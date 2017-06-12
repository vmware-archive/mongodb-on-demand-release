#!/usr/bin/env python
from subprocess import call
from sys import argv

with open("test-manifests/vms.json", "r") as myfile:
    vms = myfile.read().replace('\n', ' ').replace('\r', '')

with open("test-manifests/params.json", "r") as myfile:
    params = myfile.read().replace('\n', ' ').replace('\r', '')

with open("test-manifests/deployment.yml", "r") as myfile:
    deployment = myfile.read()

call(['go', 'run', 'main.go', argv[1], 'some-binding-id', vms,
            deployment, params])
