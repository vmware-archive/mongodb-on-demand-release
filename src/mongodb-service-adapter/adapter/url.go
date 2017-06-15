package adapter

import (
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

type DashboardURLGenerator struct {}

func (d *DashboardURLGenerator) DashboardUrl(instanceID string, plan serviceadapter.Plan, manifest bosh.BoshManifest) (serviceadapter.DashboardUrl, error) {
	// TODO: implement url generator
	return serviceadapter.DashboardUrl{DashboardUrl: "http://todo.com"}, nil
}
