#!/usr/bin/env bash
set -euo pipefail
[ 'true' = "${DEBUG:-}" ] && set -x

base=$PWD

VERSION=$(cat "$base"/version/number)
if [ -z "${VERSION:-}" ]; then
  echo "missing version number"
  exit 1
fi

cp "$base"/on-demand-service-broker-release/on-demand-service-broker-*.tgz "$base"/mongodb-on-demand-release/tile/resources
cp "$base"/syslog-migration-release/syslog-migration-*.tgz "$base"/mongodb-on-demand-release/tile/resources
cp "$base"/pcf-mongodb-helpers/pcf-mongodb-helpers-*.tgz "$base"/mongodb-on-demand-release/tile/resources

(
cd mongodb-on-demand-release

tarball_path="$base/mongodb-on-demand-release/tile/resources/mongodb-${VERSION}.tgz"
mkdir -p "$(dirname "$tarball_path")"
bosh -n create-release --sha2 --tarball="$tarball_path" --version="${VERSION}"
)

(
cd mongodb-on-demand-release/tile

yq w -i tile.yml packages.[3].path "$(ls resources/mongodb-*.tgz)"
yq w -i tile.yml packages.[3].jobs[0].properties.service_deployment.releases[0].version "${VERSION}"
tile build "${VERSION}"
)

cp "$base"/mongodb-on-demand-release/tile/product/mongodb-on-demand-*.pivotal "$base"/artifacts
