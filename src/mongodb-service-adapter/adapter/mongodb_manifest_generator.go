package adapter

import (
	"fmt"
	"log"
	"strings"

	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

const (
	StemcellAlias           = "mongodb-stemcell"
	MongodInstanceGroupName = "mongod_node"
	MongodJobName           = "mongod_node"
)

var (
	MongodJobs = []string{MongodJobName}
)

type ManifestGenerator struct {
	Logger *log.Logger
}

func (m *ManifestGenerator) logf(msg string, v ...interface{}) {
	if m.Logger != nil {
		m.Logger.Printf(msg, v...)
	}
}

func (m ManifestGenerator) GenerateManifest(
	serviceDeployment serviceadapter.ServiceDeployment,
	plan serviceadapter.Plan,
	requestParams serviceadapter.RequestParameters,
	previousManifest *bosh.BoshManifest,
	previousPlan *serviceadapter.Plan) (bosh.BoshManifest, error) {

	arbitraryParams := requestParams.ArbitraryParams()

	mongoOps := plan.Properties["mongo_ops"].(map[string]interface{})

	username := mongoOps["username"].(string)
	apiKey := mongoOps["api_key"].(string)
	url := mongoOps["url"].(string)

	oc := &OMClient{Url: url, Username: username, ApiKey: apiKey}

	adminPassword := oc.RandomString(20)

	group, err := oc.CreateGroup()
	if err != nil {
		return bosh.BoshManifest{}, fmt.Errorf("could not create new group (%s)", err.Error())
	}

	m.logf("created group %s (%s)", group.Name, group.ID)

	releases := []bosh.Release{}
	for _, release := range serviceDeployment.Releases {
		releases = append(releases, bosh.Release{
			Name:    release.Name,
			Version: release.Version,
		})
	}

	mongodInstanceGroup := findInstanceGroup(plan, MongodInstanceGroupName)
	if mongodInstanceGroup == nil {
		return bosh.BoshManifest{}, fmt.Errorf("no definition found for instance group '%s'", MongodInstanceGroupName)
	}

	mongodJobs, err := gatherJobs(serviceDeployment.Releases, MongodJobs)
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

	mongodProperties, err := mongodProperties(serviceDeployment.DeploymentName, plan.Properties, arbitraryParams, previousManifest)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	manifestProperties, err := manifestProperties(serviceDeployment.DeploymentName, group, plan.Properties, adminPassword)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	m.logf("service releases %+v", serviceDeployment.Releases)

	configAgentRelease, err := findReleaseForJob(serviceDeployment.Releases, "mongodb_config_agent")

	m.logf("conf agent releases %+v", configAgentRelease)

	configAgentProperties, err := configAgentProperties(serviceDeployment.DeploymentName,
		group, plan.Properties, adminPassword)

	if err != nil {
		return bosh.BoshManifest{}, err
	}

	return bosh.BoshManifest{
		Name:     serviceDeployment.DeploymentName,
		Releases: releases,
		Stemcells: []bosh.Stemcell{
			{
				Alias:   StemcellAlias,
				OS:      serviceDeployment.Stemcell.OS,
				Version: serviceDeployment.Stemcell.Version,
			},
		},
		InstanceGroups: []bosh.InstanceGroup{
			{
				Name:               MongodInstanceGroupName,
				Instances:          mongodInstanceGroup.Instances,
				Jobs:               mongodJobs,
				VMType:             mongodInstanceGroup.VMType,
				Stemcell:           StemcellAlias,
				PersistentDiskType: mongodInstanceGroup.PersistentDiskType,
				AZs:                mongodInstanceGroup.AZs,
				Networks:           mongodNetworks,
				Properties:         mongodProperties,
			},
			{
				Name:      "mongodb-config-agent",
				Instances: 1,
				Jobs: []bosh.Job{
					{
						Name:    "mongodb_config_agent",
						Release: configAgentRelease.Name,
						Consumes: map[string]interface{}{
							"mongod_node": bosh.ConsumesLink{From: "mongod_node"},
						},
					},
				},
				VMType:     mongodInstanceGroup.VMType,
				Stemcell:   StemcellAlias,
				AZs:        mongodInstanceGroup.AZs,
				Networks:   mongodNetworks,
				Properties: configAgentProperties,
			},
		},
		Update: bosh.Update{
			Canaries:        1,
			CanaryWatchTime: "3000-180000",
			UpdateWatchTime: "3000-180000",
			MaxInFlight:     4,
		},
		Properties: manifestProperties,
	}, nil
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

		job := bosh.Job{
			Name:    requiredJob,
			Release: release.Name,
			Provides: map[string]bosh.ProvidesLink{
				"mongod_node": bosh.ProvidesLink{As: "mongod_node"},
			},
			Consumes: map[string]interface{}{
				"mongod_node": bosh.ConsumesLink{From: "mongod_node"},
			},
		}

		jobs = append(jobs, job)
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

// TODO: figure out what's going on here
func mongodProperties(deploymentName string, planProperties serviceadapter.Properties, arbitraryParams map[string]interface{}, previousManifest *bosh.BoshManifest) (map[string]interface{}, error) {
	return map[string]interface{}{
	// "mongo_ops": mongoOps,
	// "spark_master": map[interface{}]interface{}{
	// 	"port":       SparkMasterPort,
	// 	"webui_port": SparkMasterWebUIPort,
	// },
	}, nil
}

func manifestProperties(deploymentName string, group Group, planProperties serviceadapter.Properties, adminPassword string) (map[string]interface{}, error) {
	mongoOps := planProperties["mongo_ops"].(map[string]interface{})
	url := mongoOps["url"].(string)

	return map[string]interface{}{
		"mongo_ops": map[string]string{
			"url":            url,
			"api_key":        group.AgentAPIKey,
			"group_id":       group.ID,
			"admin_password": adminPassword,
		},
	}, nil
}

func configAgentProperties(deploymentName string, group Group, planProperties serviceadapter.Properties, adminPassword string) (map[string]interface{}, error) {
	// mongo_ops.url:
	// 	description: "Mongo Ops Manager URL"
	// mongo_ops.api_key:
	//  description: "API Key for Ops Manager"
	// mongo_ops.username:
	//  description: "Username for Ops Manager"
	// mongo_ops.group_id:
	//  description: "Group Id"
	// mongo_ops.plan_id:
	//  description: "Plan identifier"

	mongoOps := planProperties["mongo_ops"].(map[string]interface{})
	url := mongoOps["url"].(string)
	username := mongoOps["username"].(string)
	apiKey := mongoOps["api_key"].(string)

	return map[string]interface{}{
		"mongo_ops": map[string]string{
			"url":            url,
			"api_key":        apiKey,
			"username":       username,
			"group_id":       group.ID,
			"plan_id":        planProperties["id"].(string),
			"admin_password": adminPassword,
		},
	}, nil
}
