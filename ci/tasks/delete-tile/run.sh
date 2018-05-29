#!/usr/bin/env bash
set -euo pipefail
[ 'true' = "${DEBUG:-}" ] && set -x

base=$PWD

PRODUCT="$(yq r $base/mongodb-on-demand-release/tile/tile.yml name)"

om="om -t $PCF_URL -u $PCF_USERNAME -p $PCF_PASSWORD -k"

echo "Retrieving current staged version of ${PRODUCT}"

product_version=$(${om} -f json deployed-products | jq -r '.[] | select(.name=="${PRODUCT}") | .version')

echo "Deleting product [${PRODUCT}], version [${product_version}] , from ${PCF_URL}"

${om} delete-product --product-name "$PRODUCT" --product-version "$product_version"
