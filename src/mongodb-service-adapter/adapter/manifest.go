package adapter

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

const (
	StemcellAlias           = "mongodb-stemcell"
	MongodInstanceGroupName = "mongod_node"
	MongodJobName           = "mongod_node"
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

	m.logf("request params: %#v", requestParams)

	arbitraryParams := requestParams.ArbitraryParams()

	mongoOps := plan.Properties["mongo_ops"].(map[string]interface{})

	username := mongoOps["username"].(string)
	apiKey := mongoOps["api_key"].(string)

	// trim trailing slash
	url := mongoOps["url"].(string)
	url = strings.TrimRight(url, "/")

	oc := &OMClient{Url: url, Username: username, ApiKey: apiKey}

	adminPassword, err := GenerateString(20)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	id, err := GenerateString(8)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	// ugly, but that was the only way to pass CC creds
	// along with each plan's parameters.
	cfOps := plan.Properties["cf"].(map[string]interface{})

	c := &cfclient.Config{
		ApiAddress:        cfOps["url"].(string),
		Username:          cfOps["username"].(string),
		Password:          cfOps["password"].(string),
		SkipSslValidation: cfOps["disable_ssl_cert_verification"].(bool),
	}

	cc, err := cfclient.NewClient(c)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	si, err := cc.ServiceInstanceByGuid(requestParams["service_id"].(string))
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	group, err := oc.CreateGroup(si.Name + "_" + id)
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

	mongodJobs, err := gatherJobs(serviceDeployment.Releases, []string{MongodJobName})
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

	configAgentRelease, err := findReleaseForJob(serviceDeployment.Releases, "mongodb_config_agent")
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	engineVersion, ok := arbitraryParams["version"].(string)
	if engineVersion == "" || !ok {
		engineVersion = "3.2.7" // TODO: make it configurable in deployment manifest
	}

	// sharded_cluster parameters
	replicas := 0
	routers := 0
	configServers := 0

	// total number of instances
	//
	// standalone:      always one
	// replica_set:     number of replicas
	// sharded_cluster: shards*replicas + config_servers + mongos
	instances := mongodInstanceGroup.Instances

	planID := plan.Properties["id"].(string)
	switch planID {
	case PlanStandalone:
		// ok
	case PlanReplicaSet:
		if r, ok := arbitraryParams["replicas"].(float64); ok && r > 0 {
			instances = int(r)
		}
	case PlanShardedCluster:
		shards := 2
		if s, ok := arbitraryParams["shards"].(float64); ok && s > 0 {
			shards = int(s)
		}

		replicas = 2
		if r, ok := arbitraryParams["replicas"].(float64); ok && r > 0 {
			replicas = int(r)
		}

		configServers = 2
		if c, ok := arbitraryParams["config_servers"].(float64); ok && c > 0 {
			configServers = int(c)
		}

		routers = 2
		if r, ok := arbitraryParams["mongos"].(float64); ok && r > 0 {
			routers = int(r)
		}

		instances = routers + configServers + shards*replicas
	default:
		return bosh.BoshManifest{}, fmt.Errorf("unknown plan: %s", planID)
	}

	manifest := bosh.BoshManifest{
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
				Instances:          instances,
				Jobs:               mongodJobs,
				VMType:             mongodInstanceGroup.VMType,
				Stemcell:           StemcellAlias,
				PersistentDiskType: mongodInstanceGroup.PersistentDiskType,
				AZs:                mongodInstanceGroup.AZs,
				Networks:           mongodNetworks,
				Properties:         map[string]interface{}{},
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
				VMType:   mongodInstanceGroup.VMType,
				Stemcell: StemcellAlias,
				AZs:      mongodInstanceGroup.AZs,
				Networks: mongodNetworks,

				// See mongodb_config_agent job spec
				Properties: map[string]interface{}{
					"mongo_ops": map[string]interface{}{
						"id":             id,
						"url":            url,
						"api_key":        apiKey,
						"username":       username,
						"group_id":       group.ID,
						"plan_id":        planID,
						"admin_password": adminPassword,
						"engine_version": engineVersion,
						"routers":        routers,
						"config_servers": configServers,
						"replicas":       replicas,
					},
				},
			},
		},
		Update: bosh.Update{
			Canaries:        1,
			CanaryWatchTime: "3000-180000",
			UpdateWatchTime: "3000-180000",
			MaxInFlight:     4,
		},
		Properties: map[string]interface{}{
			"mongo_ops": map[string]interface{}{
				"url":            url,
				"api_key":        group.AgentAPIKey,
				"group_id":       group.ID,
				"admin_password": adminPassword,

				// options needed for binding
				"plan_id":        planID,
				"routers":        routers,
				"config_servers": configServers,
				"replicas":       replicas,
			},
		},
	}

	m.logf("generated manifest: %#v", manifest)
	return manifest, nil
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
				"mongod_node": {As: "mongod_node"},
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
