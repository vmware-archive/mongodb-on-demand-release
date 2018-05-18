#!/usr/bin/env bash
set -euo pipefail
[ 'true' = "${DEBUG:-}" ] && set -x

base=$PWD

VERSION=$(cat "$base"/version/number)
if [ -z "${VERSION:-}" ]; then
  echo "missing version number"
  exit 1
fi

cp "$base"/on-demand-service-broker/on-demand-service-broker-*.tgz "$base"/mongodb-on-demand-release/tile/resources
cp "$base"/pcf-mongodb-helpers/pcf-mongodb-helpers-*.tgz "$base"/mongodb-on-demand-release/tile/resources

(
cd mongodb-on-demand-release

tarball_path="$base/mongodb-on-demand-release/tile/resources/mongodb-${VERSION}.tgz"
mkdir -p "$(dirname "$tarball_path")"
bosh -n create-release --sha2 --tarball="$tarball_path" --version="${VERSION}"
)

(
cd mongodb-on-demand-release/tile

yq w -i tile.yml packages.[2].path "$(ls resources/mongodb-*.tgz)"
yq w -i tile.yml packages.[2].jobs[0].properties.service_deployment.releases[0].version "${VERSION}"
tile build "${VERSION}"
)

cp "$base"/mongodb-on-demand-release/tile/product/mongodb-on-demand-*.pivotal "$base"/release

TILE_FILE=`cd release; ls *-${VERSION}.pivotal`
if [ -z "${TILE_FILE}" ]; then
	echo "No files matching release/*.pivotal"
	ls -lR release
	exit 1
fi
SHA256=$(sha256sum release/"$TILE_FILE" | cut -d ' ' -f 1)
cat > "$base"/release/notification <<EOF
<!here> New build v${VERSION} released!
Build SHA256: ${SHA256}
EOF
