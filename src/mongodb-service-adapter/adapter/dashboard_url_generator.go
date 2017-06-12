package adapter

import "github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
import "github.com/pivotal-cf/on-demand-services-sdk/bosh"

type DashboardURLGenerator struct {
}

func (d *DashboardURLGenerator) DashboardUrl(instanceID string, plan serviceadapter.Plan, manifest bosh.BoshManifest) (serviceadapter.DashboardUrl, error) {
	return serviceadapter.DashboardUrl{DashboardUrl: "http://todo.com"}, nil
}
