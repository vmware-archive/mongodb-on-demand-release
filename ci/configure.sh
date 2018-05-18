#!/usr/bin/env bash

set -eu

fly -t altoros set-pipeline -n \
 -p mongodb-on-demand-release \
 -c ./pipeline.yml \
 -l <(lpass show --note "pcf:mongodb-on-demand-release")
