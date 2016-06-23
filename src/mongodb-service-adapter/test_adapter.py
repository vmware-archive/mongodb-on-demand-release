#!/usr/bin/env python
from subprocess import call
from sys import argv

with open("test-manifests/bosh-info.json", "r") as myfile:
    bosh_info = myfile.read().replace('\n', ' ').replace('\r', '')

with open("test-manifests/params.json", "r") as myfile:
    params = myfile.read().replace('\n', ' ').replace('\r', '')

with open("test-manifests/plan.json", "r") as myfile:
    plan = myfile.read().replace('\n', ' ').replace('\r', '')

with open("test-manifests/service-releases.json", "r") as myfile:
    service_releases = myfile.read().replace('\n', ' ').replace('\r', '')

call(['go', 'run', 'cmd/mongodb-service-adapter/main.go', argv[1], bosh_info, plan, params, '---', '{}'])
