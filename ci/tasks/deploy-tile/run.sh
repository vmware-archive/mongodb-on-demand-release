#!/usr/bin/env bash
set -euo pipefail
[ 'true' = "${DEBUG:-}" ] && set -x

base=$PWD

VERSION=$(cat "$base"/version/number)
if [ -z "${VERSION:-}" ]; then
  echo "missing version number"
  exit 1
fi

TILE_FILE=`cd artifacts; ls *-${VERSION}.pivotal`
if [ -z "${TILE_FILE}" ]; then
	echo "No files matching artifacts/*.pivotal"
	ls -lR artifacts
	exit 1
fi

STEMCELL_FILE=`cd stemcell; ls *bosh-stemcell-*.tgz`
if [ -z "${STEMCELL_FILE}" ]; then
	echo "No files matching stemcell/*.tgz"
	ls -lR stemcell
	exit 1
fi

PRODUCT="$(cat $base/mongodb-on-demand-release/tile/tile.yml | grep '^name' | cut -d' ' -f 2)"

om="om -t $PCF_URL -u $PCF_USERNAME -p $PCF_PASSWORD -k"

${om} upload-product --product "artifacts/$TILE_FILE"
${om} upload-stemcell --stemcell "stemcell/$STEMCELL_FILE"
${om} available-products
${om} stage-product --product-name "$PRODUCT" --product-version "$VERSION"

cat > network-config.json << EOF
{
  "singleton_availability_zone": {
    "name": "$SINGELTON_AZ"
  },
  "other_availability_zones": [
    {
      "name": "$OTHER_AZ"
    }
  ],
  "network": {
    "name": "$PCF_NETWORK"
  },
  "service_network": {
    "name": "${SERVICE_NETWORK}"
  }
}
EOF

cat > missing-properties.json << EOF
{
  ".properties.url": {
    "value": "${MONGO_OPS_URL}"
  },
  ".properties.username": {
    "value": "${MONGO_OPS_USERNAME}"
  },
  ".properties.api_key": {
    "value": {
      "secret": "${MONGO_OPS_API_KEY}"
    }
  },
  ".properties.az_multi_select": {
    "value": ["${AZ}"]
  },
  ".properties.ca_cert": {
    "value": "${MONGO_CA_CERT}"
  },
  ".properties.rsa_certificate": {
    "value": {
      "cert_pem":"${MONGO_OPS_RSA_CERT}",
      "private_key_pem": "${MONGO_OPS_RSA_KEY}"
    }
  }
}
EOF

${om} configure-product --product-name "$PRODUCT" --product-network "$(cat network-config.json| jq -cr .)" --product-properties "$(cat missing-properties.json| jq -cr .)"

STAGED=$(${om} curl --path /api/v0/staged/products)
RESULT=$(echo "$STAGED" | jq --arg product_name "$PRODUCT" 'map(select(.type == $product_name)) | .[].guid')
DATA=$(echo '{"deploy_products": []}' | jq ".deploy_products += [$RESULT]")

${om} curl --path /api/v0/installations --request POST --data "$DATA"
${om} apply-changes --skip-deploy-products="true"
${om} delete-unused-products
