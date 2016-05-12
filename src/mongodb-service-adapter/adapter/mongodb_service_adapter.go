package adapter

import (
	"fmt"
	"strings"
	"log"
	"os"

	"github.com/pivotal-cf/on-demand-service-broker-sdk/bosh"
	"github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"
)

const (
	StemcellAlias           = "mongodb-stemcell"
	MongodInstanceGroupName = "mongod_node"
	MongodJobName           = "mongod_node"
)

var (
	MongodJobs = []string{MongodJobName}
)

type Adapter struct { }

func (a Adapter) GenerateManifest(
	boshInfo serviceadapter.BoshInfo,
	serviceReleases serviceadapter.ServiceReleases,
	plan serviceadapter.Plan,
	arbitraryParams map[string]interface{},
	previousManifest *bosh.BoshManifest,
) (bosh.BoshManifest, error) {

	logger := log.New(os.Stderr, "[mongodb-service-adapter] ", log.LstdFlags)

	// doc := oc.LoadDoc(plan.Properties["id"].(string))

	mongoOps := plan.Properties["mongo_ops"].(map[string]interface{})

	username := mongoOps["username"].(string)
	apiKey := mongoOps["api_key"].(string)
	address := mongoOps["address"].(string)
	scheme := mongoOps["scheme"].(string)
	port := mongoOps["port"].(float64)

	url := fmt.Sprintf("%s://%s:%g", scheme, address, port)

	oc := &OMClient { Url: url, Username: username, ApiKey: apiKey }

	group, _ := oc.CreateGroup()
	logger.Printf("created group %s (%s)", group.Name, group.ID)
	// url, groupID, apiKey := oc.CreateCluster()

	releases := []bosh.Release{}
	for _, release := range serviceReleases {
		releases = append(releases, bosh.Release{
			Name:    release.Name,
			Version: release.Version,
		})
	}

	mongodInstanceGroup := findInstanceGroup(plan, MongodInstanceGroupName)
	if mongodInstanceGroup == nil {
		return bosh.BoshManifest{}, fmt.Errorf("no definition found for instance group '%s'", MongodInstanceGroupName)
	}

	mongodJobs, err := gatherJobs(serviceReleases, MongodJobs)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	mongodNetworks := []bosh.Network{}
	for _, network := range mongodInstanceGroup.Networks {
		mongodNetworks = append(mongodNetworks, bosh.Network{Name: network})
	}
	if len(mongodNetworks) == 0 {
		return bosh.BoshManifest{}, fmt.Errorf("no networks definition found for instance group '%s'", MongodInstanceGroupName)
	}

	mongodProperties, err := mongodProperties(boshInfo.Name, plan.Properties, arbitraryParams, previousManifest)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	return bosh.BoshManifest{
		Name:     boshInfo.Name,
		Releases: releases,
		Stemcells: []bosh.Stemcell{
			{
				Alias:   StemcellAlias,
				OS:      boshInfo.StemcellOS,
				Version: boshInfo.StemcellVersion,
			},
		},
		InstanceGroups: []bosh.InstanceGroup{
			{
				Name:               MongodInstanceGroupName,
				Instances:          mongodInstanceGroup.Instances,
				Jobs:               mongodJobs,
				VMType:             mongodInstanceGroup.VMType,
				Stemcell:           StemcellAlias,
				PersistentDiskType: mongodInstanceGroup.PersistentDisk,
				AZs:                mongodInstanceGroup.AZs,
				Networks:           mongodNetworks,
				Properties:         mongodProperties,
			},
		},
		Update: bosh.Update{
			Canaries:        1,
			CanaryWatchTime: "3000-180000",
			UpdateWatchTime: "3000-180000",
			MaxInFlight:     4,
		},
		Properties: map[string]interface{}{},
	}, nil
}

func (a Adapter) CreateBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest) (map[string]interface{}, error) {

	return map[string]interface{}{}, nil
}

func (a Adapter) DeleteBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest) error {
	return nil
}

func findInstanceGroup(plan serviceadapter.Plan, jobName string) *serviceadapter.InstanceGroup {
	for _, instanceGroup := range plan.InstanceGroups {
		if instanceGroup.Name == jobName {
			return &instanceGroup
		}
	}
	return nil
}

func gatherJobs(releases serviceadapter.ServiceReleases, requiredJobs []string) ([]bosh.Job, error) {
	jobs := []bosh.Job{}

	for _, requiredJob := range requiredJobs {
		release, err := findReleaseForJob(releases, requiredJob)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, bosh.Job{Name: requiredJob, Release: release.Name})
	}

	return jobs, nil
}

func findReleaseForJob(releases serviceadapter.ServiceReleases, requiredJob string) (serviceadapter.ServiceRelease, error) {
	releasesThatProvideRequiredJob := serviceadapter.ServiceReleases{}

	for _, release := range releases {
		for _, providedJob := range release.Jobs {
			if providedJob == requiredJob {
				releasesThatProvideRequiredJob = append(releasesThatProvideRequiredJob, release)
			}
		}
	}

	if len(releasesThatProvideRequiredJob) == 0 {
		return serviceadapter.ServiceRelease{}, fmt.Errorf("no release provided for job '%s'", requiredJob)
	}

	if len(releasesThatProvideRequiredJob) > 1 {
		releaseNames := []string{}
		for _, release := range releasesThatProvideRequiredJob {
			releaseNames = append(releaseNames, release.Name)
		}

		return serviceadapter.ServiceRelease{}, fmt.Errorf("job '%s' defined in multiple releases: %s", requiredJob, strings.Join(releaseNames, ", "))
	}

	return releasesThatProvideRequiredJob[0], nil
}

func mongodProperties(deploymentName string, planProperties serviceadapter.Properties, arbitraryParams map[string]interface{}, previousManifest *bosh.BoshManifest) (map[string]interface{}, error) {
	return map[string]interface{}{
		// "spark_master": map[interface{}]interface{}{
		// 	"port":       SparkMasterPort,
		// 	"webui_port": SparkMasterWebUIPort,
		// },
	}, nil
}
