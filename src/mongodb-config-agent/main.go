package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/cf-platform-eng/mongodb-on-demand-release/src/mongodb-service-adapter/adapter"
)

var (
	configFilePath string
)

func main() {
	flag.StringVar(&configFilePath, "config", "", "Location of the config file")
	flag.Parse()

	config, err := LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	logger := log.New(os.Stderr, "[mongodb-config-agent] ", log.LstdFlags)
	omClient := adapter.OMClient{Url: config.URL, Username: config.Username, ApiKey: config.APIKey}

	nodes := strings.Split(config.NodeAddresses, ",")
	ctx := &adapter.DocContext{
		ID:            config.ID,
		Key:           "GrSLAAsHGXmJOrvElJ2AHTGauvH4O0EFT1r8byvb0G9sTU0viVX21PwUMqBjyXB9WrZP9QvEmCQIF1wOqJofyWmx7wWZqpO69dnc9GUWcpGQLr7eVyKTs99WAPXR3kXpF4MVrHdBMEDfRfhytgomgAso96urN6eC8RaUpjX4Bf9HcAEJwfddZshin97XKJDmqCaqAfORNnf1e8hkfTIwYg1tvIpwemmEF4TkmOgK09N5dINyejyWMU8iWG8FqW5MfQ8A2DrtIdyGSKLH05s7H1dXyADjDECaC77QqLXTx7gWyHca3I0K92PuVFoOs5385vzqTYN3kVgFotSdXgoM8Zt5QIoj2lX4PYqm2TWsVp0s15JELikH8bNVIIMGiSSWJEWGU1PVEXD7V7cYepDb88korMjr3wbh6kZ76Q7F2RtfJqkd4hKw7B5OCX04b5eppkjL598iCpSUUx3X9C6fFavWj2DrHsv9DY86iCWBlcG08DRPKs9EPizCW4jNZtJcm3T7WlcI0MZMKOtsKOCWBZA0C9YnttNrp4eTsQ1U43StiIRPqp2K8rrQAu6etURH0RHedazHeeukTWI7iTG1dZpYk9EyittZ72qKXLNLhi5vJ9TlYw8O91vihB1nJwwA3B1WbiYhkqqRzoL0cQpXJMUsUlsoSP6Q70IMU92vEHbUmna5krESPLeJfQBKGQPNVVE63XYBh2TnvFTdi6koitu209wMFUnHZrzWj3UWGqsyTqqHbPl4RhRLFe24seRwV2SbUuLygBIdptKHnA3kutAbHzsWTT8UxOaiQzFV4auxounrgXj7MoMWEVKKS8AHkELPILGqFVFC8BZsfPC0WacSN5Rg5SaCvfs74hcsCQ3ghq9PyxEb2fbHUiaCjnsBcXqzQw9AjZJG4yX0ubEwicP0bKB6y3w4PUQqdouxH5y16OgkUjrZgodJfRLgP9vqGbHNDpj4yBuswluvCFBh38gBoSIQu11qtQmk43n4G8Dskn0DrJ32l2Gz35q5LaKT",
		AdminPassword: config.AdminPassword,
		Nodes:         nodes,
		Version:       config.EngineVersion,
	}

	if config.PlanID == adapter.PlanShardedCluster {
		var err error
		ctx.Cluster, err = adapter.NodesToCluster(nodes, config.Routers, config.ConfigServers, config.Replicas)
		if err != nil {
			logger.Fatal(err)
		}
	}

	logger.Printf("%+v", nodes)
	doc, err := omClient.LoadDoc(config.PlanID, ctx)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println(doc)

	for {
		logger.Printf("Checking group %s", config.GroupID)

		groupHosts, err := omClient.GetGroupHosts(config.GroupID)
		if err != nil {
			logger.Fatal(err)
		}

		//	logger.Printf("total number of hosts *** %v", groupHosts.TotalCount)
		if groupHosts.TotalCount == 0 {
			logger.Printf("Host count for %s is 0, configuring...", config.GroupID)

			err = omClient.ConfigureGroup(doc, config.GroupID)
			if err != nil {
				logger.Fatal(err)
			}

			logger.Printf("Configured group %s", config.GroupID)
		}

		time.Sleep(30 * time.Second)
	}
}
