package adapter

import (
	"fmt"

	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

type DashboardURLGenerator struct{}

func (d *DashboardURLGenerator) DashboardUrl(instanceID string, plan serviceadapter.Plan, manifest bosh.BoshManifest) (serviceadapter.DashboardUrl, error) {
	properties := manifest.Properties["mongo_ops"].(map[interface{}]interface{})
	url := properties["url"].(string)
	groupID := properties["group_id"].(string)

	return serviceadapter.DashboardUrl{
		DashboardUrl: fmt.Sprintf("%s/v2/%s", url, groupID),
	}, nil
}
