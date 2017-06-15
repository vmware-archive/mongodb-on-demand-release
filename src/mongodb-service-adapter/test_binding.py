#!/usr/bin/env python
from subprocess import call
from sys import argv

with open("testdata/vms.json", "r") as myfile:
    vms = myfile.read().replace('\n', ' ').replace('\r', '')

with open("testdata/params.json", "r") as myfile:
    params = myfile.read().replace('\n', ' ').replace('\r', '')

with open("testdata/deployment.yml", "r") as myfile:
    deployment = myfile.read()

call(['go', 'run', 'main.go', argv[1], 'some-binding-id', vms,
            deployment, params])
