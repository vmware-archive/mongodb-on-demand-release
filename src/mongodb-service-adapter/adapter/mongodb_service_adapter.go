package adapter

import (
	"fmt"
	"strings"

	"github.com/pivotal-cf/on-demand-service-broker-sdk/bosh"
	"github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"
)


type Adapter struct {
}

func (a Adapter) GenerateManifest(
	boshInfo serviceadapter.BoshInfo,
	serviceReleases serviceadapter.ServiceReleases,
	plan serviceadapter.Plan,
	arbitraryParams map[string]interface{},
	previousManifest *bosh.BoshManifest,
) (bosh.BoshManifest, error) {

	return bosh.BoshManifest{}, nil
}

func (a Adapter) CreateBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest) (map[string]interface{}, error) {

	return map[string]interface{} {}, nil
}

func (a Adapter) DeleteBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest) error {
	return nil
}
