package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/cf-platform-eng/mongodb-on-demand-release/src/mongodb-service-adapter/adapter"
)

// TODO: pass json instead of flags
var (
	id            string
	url           string
	username      string
	apiKey        string
	groupID       string
	planID        string
	nodeAddresses string
	adminPassword string
	engineVersion string
	replicas      int
)

func main() {
	flag.StringVar(&id, "id", "", "ID")
	flag.StringVar(&url, "url", "", "MOM URL")
	flag.StringVar(&username, "username", "", "MOM Username")
	flag.StringVar(&apiKey, "api-key", "", "MOM API Key")
	flag.StringVar(&groupID, "group", "", "MOM Group ID")
	flag.StringVar(&planID, "plan", "", "The name of the service plan")
	flag.StringVar(&nodeAddresses, "nodes", "", "Comma separated list of addresses")
	flag.StringVar(&adminPassword, "admin-password", "", "Admin password for the mongo instance")
	flag.StringVar(&engineVersion, "engine-version", "", "Engine version")
	flag.IntVar(&replicas, "replicas", 0, "replicas per shard")
	flag.Parse()

	logger := log.New(os.Stderr, "[mongodb-config-agent] ", log.LstdFlags)
	omClient := adapter.OMClient{Url: url, Username: username, ApiKey: apiKey}

	nodes := strings.Split(nodeAddresses, ",")
	ctx := &adapter.DocContext{
		ID:            id,
		Key:           "GrSLAAsHGXmJOrvElJ2AHTGauvH4O0EFT1r8byvb0G9sTU0viVX21PwUMqBjyXB9WrZP9QvEmCQIF1wOqJofyWmx7wWZqpO69dnc9GUWcpGQLr7eVyKTs99WAPXR3kXpF4MVrHdBMEDfRfhytgomgAso96urN6eC8RaUpjX4Bf9HcAEJwfddZshin97XKJDmqCaqAfORNnf1e8hkfTIwYg1tvIpwemmEF4TkmOgK09N5dINyejyWMU8iWG8FqW5MfQ8A2DrtIdyGSKLH05s7H1dXyADjDECaC77QqLXTx7gWyHca3I0K92PuVFoOs5385vzqTYN3kVgFotSdXgoM8Zt5QIoj2lX4PYqm2TWsVp0s15JELikH8bNVIIMGiSSWJEWGU1PVEXD7V7cYepDb88korMjr3wbh6kZ76Q7F2RtfJqkd4hKw7B5OCX04b5eppkjL598iCpSUUx3X9C6fFavWj2DrHsv9DY86iCWBlcG08DRPKs9EPizCW4jNZtJcm3T7WlcI0MZMKOtsKOCWBZA0C9YnttNrp4eTsQ1U43StiIRPqp2K8rrQAu6etURH0RHedazHeeukTWI7iTG1dZpYk9EyittZ72qKXLNLhi5vJ9TlYw8O91vihB1nJwwA3B1WbiYhkqqRzoL0cQpXJMUsUlsoSP6Q70IMU92vEHbUmna5krESPLeJfQBKGQPNVVE63XYBh2TnvFTdi6koitu209wMFUnHZrzWj3UWGqsyTqqHbPl4RhRLFe24seRwV2SbUuLygBIdptKHnA3kutAbHzsWTT8UxOaiQzFV4auxounrgXj7MoMWEVKKS8AHkELPILGqFVFC8BZsfPC0WacSN5Rg5SaCvfs74hcsCQ3ghq9PyxEb2fbHUiaCjnsBcXqzQw9AjZJG4yX0ubEwicP0bKB6y3w4PUQqdouxH5y16OgkUjrZgodJfRLgP9vqGbHNDpj4yBuswluvCFBh38gBoSIQu11qtQmk43n4G8Dskn0DrJ32l2Gz35q5LaKT",
		AdminPassword: adminPassword,
		Nodes:         nodes,
		Version:       engineVersion,
	}

	if planID == adapter.PlanShardedSet {
		var err error
		ctx.Cluster, err = NodesToCluster(nodes, replicas, replicas, replicas)
		if err != nil {
			logger.Fatal(err)
		}
	}

	logger.Printf("%+v", nodes)
	doc, err := omClient.LoadDoc(adapter.Plan(planID), ctx)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println(doc)

	for {
		logger.Printf("Checking group %s", groupID)

		groupHosts, err := omClient.GetGroupHosts(groupID)
		if err != nil {
			logger.Fatal(err)
		}

		//	logger.Printf("total number of hosts *** %v", groupHosts.TotalCount)
		if groupHosts.TotalCount == 0 {
			logger.Printf("Host count for %s is 0, configuring...", groupID)

			err = omClient.ConfigureGroup(doc, groupID)
			if err != nil {
				logger.Fatal(err)
			}

			logger.Printf("Configured group %s", groupID)
		}

		time.Sleep(30 * time.Second)
	}
}

// TODO: validate input
func NodesToCluster(nodes []string, routers, configServers, replicas int) (*adapter.Cluster, error) {
	c := &adapter.Cluster{
		Routers:       nodes[:routers],
		ConfigServers: nodes[routers : routers+configServers],
	}

	nodes = nodes[routers+configServers:]
	c.Shards = make([][]string, 0, len(nodes)/replicas)
	for i := 0; i < len(nodes)/replicas; i++ {
		c.Shards = append(c.Shards, make([]string, 0, replicas))
		for j := 0; j < replicas; j++ {
			c.Shards[i] = append(c.Shards[i], nodes[i*replicas+j])
		}
	}
	return c, nil
}
