#!/usr/bin/env python
from subprocess import call
from sys import argv

with open("testdata/bosh-info.json", "r") as myfile:
    bosh_info = myfile.read().replace('\n', ' ').replace('\r', '')

with open("testdata/params.json", "r") as myfile:
    params = myfile.read().replace('\n', ' ').replace('\r', '')

with open("testdata/plan.json", "r") as myfile:
    plan = myfile.read().replace('\n', ' ').replace('\r', '')

with open("testdata/service-releases.json", "r") as myfile:
    service_releases = myfile.read().replace('\n', ' ').replace('\r', '')

call(['go', 'run', 'main.go', argv[1], bosh_info, plan, params, '---', '{}'])
