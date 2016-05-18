package adapter

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/mgo.v2"

	// "gopkg.in/mgo.v2/bson"
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

type Adapter struct{}

func (a Adapter) GenerateManifest(
	boshInfo serviceadapter.BoshInfo,
	serviceReleases serviceadapter.ServiceReleases,
	plan serviceadapter.Plan,
	arbitraryParams map[string]interface{},
	previousManifest *bosh.BoshManifest,
) (bosh.BoshManifest, error) {

	logger := log.New(os.Stderr, "[mongodb-service-adapter] ", log.LstdFlags)

	mongoOps := plan.Properties["mongo_ops"].(map[string]interface{})

	username := mongoOps["username"].(string)
	apiKey := mongoOps["api_key"].(string)
	url := mongoOps["url"].(string)

	oc := &OMClient{Url: url, Username: username, ApiKey: apiKey}

	group, err := oc.CreateGroup()
	if err != nil {
		return bosh.BoshManifest{}, fmt.Errorf("could not create new group (%s)", err.Error())
	}

	logger.Printf("created group %s (%s)", group.Name, group.ID)

	// generate context
	ctx := map[string]string{
		"auto_user":      "mms-automation",
		"auto_password":  "Sy4oBX9ei0amvupBAN8lVQhj",
		"key":            "GrSLAAsHGXmJOrvElJ2AHTGauvH4O0EFT1r8byvb0G9sTU0viVX21PwUMqBjyXB9WrZP9QvEmCQIF1wOqJofyWmx7wWZqpO69dnc9GUWcpGQLr7eVyKTs99WAPXR3kXpF4MVrHdBMEDfRfhytgomgAso96urN6eC8RaUpjX4Bf9HcAEJwfddZshin97XKJDmqCaqAfORNnf1e8hkfTIwYg1tvIpwemmEF4TkmOgK09N5dINyejyWMU8iWG8FqW5MfQ8A2DrtIdyGSKLH05s7H1dXyADjDECaC77QqLXTx7gWyHca3I0K92PuVFoOs5385vzqTYN3kVgFotSdXgoM8Zt5QIoj2lX4PYqm2TWsVp0s15JELikH8bNVIIMGiSSWJEWGU1PVEXD7V7cYepDb88korMjr3wbh6kZ76Q7F2RtfJqkd4hKw7B5OCX04b5eppkjL598iCpSUUx3X9C6fFavWj2DrHsv9DY86iCWBlcG08DRPKs9EPizCW4jNZtJcm3T7WlcI0MZMKOtsKOCWBZA0C9YnttNrp4eTsQ1U43StiIRPqp2K8rrQAu6etURH0RHedazHeeukTWI7iTG1dZpYk9EyittZ72qKXLNLhi5vJ9TlYw8O91vihB1nJwwA3B1WbiYhkqqRzoL0cQpXJMUsUlsoSP6Q70IMU92vEHbUmna5krESPLeJfQBKGQPNVVE63XYBh2TnvFTdi6koitu209wMFUnHZrzWj3UWGqsyTqqHbPl4RhRLFe24seRwV2SbUuLygBIdptKHnA3kutAbHzsWTT8UxOaiQzFV4auxounrgXj7MoMWEVKKS8AHkELPILGqFVFC8BZsfPC0WacSN5Rg5SaCvfs74hcsCQ3ghq9PyxEb2fbHUiaCjnsBcXqzQw9AjZJG4yX0ubEwicP0bKB6y3w4PUQqdouxH5y16OgkUjrZgodJfRLgP9vqGbHNDpj4yBuswluvCFBh38gBoSIQu11qtQmk43n4G8Dskn0DrJ32l2Gz35q5LaKT",
		"admin_password": "password",
	}

	doc, err := oc.LoadDoc(plan.Properties["id"].(string), ctx)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	logger.Println(doc)

	err = oc.ConfigureGroup(doc, group.ID)

	logger.Printf("configured group %s (%s)", group.Name, group.ID)

	if err != nil {
		return bosh.BoshManifest{}, fmt.Errorf("could not configure group '%s' (%s)", group.Name, err.Error())
	}

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

	manifestProperties, err := manifestProperties(boshInfo.Name, group, plan.Properties, ctx["admin_password"])
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
		Properties: manifestProperties,
	}, nil
}

func (a Adapter) CreateBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest) (map[string]interface{}, error) {

	// create an admin level user
	username := fmt.Sprintf("pcf_%v", encodeID(bindingID))
	password := OMClient{}.RandomString(32)

	properties := manifest.Properties["mongo_ops"].(map[interface{}]interface{})
	adminPassword := properties["admin_password"].(string)

	servers := make([]string, len(deploymentTopology["mongod_node"]))
	for i, node := range deploymentTopology["mongod_node"] {
		servers[i] = fmt.Sprintf("%s:28000", node)
	}

	dialInfo := &mgo.DialInfo{
		Addrs:     servers,
		Username:  "admin",
		Password:  adminPassword,
		Mechanism: "SCRAM-SHA-1",
		Database:  "admin",
		FailFast:  true,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	adminDB := session.DB("admin")

	// add user to admin database with admin priveleges
	user := &mgo.User{
		Username: username,
		Password: password,
		Roles: []mgo.Role{
			mgo.RoleUserAdmin,
			mgo.RoleDBAdmin,
			mgo.RoleReadWrite,
		},
		OtherDBRoles: map[string][]mgo.Role{
			username: []mgo.Role{
				mgo.RoleUserAdmin,
				mgo.RoleDBAdmin,
				mgo.RoleReadWrite,
			},
		},
	}
	adminDB.UpsertUser(user)

	return map[string]interface{}{
		"username": username,
		"password": password,
		"database": username,
	}, nil
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

		job := bosh.Job{
			Name:    requiredJob,
			Release: release.Name,
			Provides: map[string]bosh.ProvidesLink{
				"mongod_node": bosh.ProvidesLink{As: "mongod_node"},
			},
			Consumes: map[string]bosh.ConsumesLink{
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

func encodeID(id string) string {
	b64 := base64.StdEncoding.EncodeToString([]byte(id))
	md5 := md5.Sum([]byte(b64))
	return fmt.Sprintf("%x", md5)
}
