package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/cf-platform-eng/mongodb-on-demand-release/src/mongodb-service-adapter/adapter"
)

type ConfigAgent struct{}

func partitionNodes(values []string, parts int) [][]string {

	b := [][]string{}

	for i := 0; i < parts; i++ {
		b = append(b, []string{})
		b[i] = values[i*parts : (i*parts)+parts]
	}

	return b
}

func (c ConfigAgent) PollAndConfigureGroup(url string, username string, apiKey string,
	groupID string, planID string, nodeAddresses string, adminPassword string) {

	logger := log.New(os.Stderr, "[mongodb-config-agent] ", log.LstdFlags)

	omClient := adapter.OMClient{Url: url, Username: username, ApiKey: apiKey}

	nodes := strings.Split(nodeAddresses, ",")
	partitionedNodes := [][]string{}
	ctx := map[string]interface{}{}

	if planID == "sharded_set" {
			partitionedNodes = partitionNodes(nodes, 3)
			logger.Printf("%#v", partitionedNodes)
			ctx = map[string]interface{}{
				"auto_user":        "mms-automation",
				"auto_password":    "Sy4oBX9ei0amvupBAN8lVQhj",
				"key":              "GrSLAAsHGXmJOrvElJ2AHTGauvH4O0EFT1r8byvb0G9sTU0viVX21PwUMqBjyXB9WrZP9QvEmCQIF1wOqJofyWmx7wWZqpO69dnc9GUWcpGQLr7eVyKTs99WAPXR3kXpF4MVrHdBMEDfRfhytgomgAso96urN6eC8RaUpjX4Bf9HcAEJwfddZshin97XKJDmqCaqAfORNnf1e8hkfTIwYg1tvIpwemmEF4TkmOgK09N5dINyejyWMU8iWG8FqW5MfQ8A2DrtIdyGSKLH05s7H1dXyADjDECaC77QqLXTx7gWyHca3I0K92PuVFoOs5385vzqTYN3kVgFotSdXgoM8Zt5QIoj2lX4PYqm2TWsVp0s15JELikH8bNVIIMGiSSWJEWGU1PVEXD7V7cYepDb88korMjr3wbh6kZ76Q7F2RtfJqkd4hKw7B5OCX04b5eppkjL598iCpSUUx3X9C6fFavWj2DrHsv9DY86iCWBlcG08DRPKs9EPizCW4jNZtJcm3T7WlcI0MZMKOtsKOCWBZA0C9YnttNrp4eTsQ1U43StiIRPqp2K8rrQAu6etURH0RHedazHeeukTWI7iTG1dZpYk9EyittZ72qKXLNLhi5vJ9TlYw8O91vihB1nJwwA3B1WbiYhkqqRzoL0cQpXJMUsUlsoSP6Q70IMU92vEHbUmna5krESPLeJfQBKGQPNVVE63XYBh2TnvFTdi6koitu209wMFUnHZrzWj3UWGqsyTqqHbPl4RhRLFe24seRwV2SbUuLygBIdptKHnA3kutAbHzsWTT8UxOaiQzFV4auxounrgXj7MoMWEVKKS8AHkELPILGqFVFC8BZsfPC0WacSN5Rg5SaCvfs74hcsCQ3ghq9PyxEb2fbHUiaCjnsBcXqzQw9AjZJG4yX0ubEwicP0bKB6y3w4PUQqdouxH5y16OgkUjrZgodJfRLgP9vqGbHNDpj4yBuswluvCFBh38gBoSIQu11qtQmk43n4G8Dskn0DrJ32l2Gz35q5LaKT",
				"admin_password":   adminPassword,
				"nodes":            nodes,
				"partitionedNodes": partitionedNodes,
			}
  } else {
			ctx = map[string]interface{}{
				"auto_user":        "mms-automation",
				"auto_password":    "Sy4oBX9ei0amvupBAN8lVQhj",
				"key":              "GrSLAAsHGXmJOrvElJ2AHTGauvH4O0EFT1r8byvb0G9sTU0viVX21PwUMqBjyXB9WrZP9QvEmCQIF1wOqJofyWmx7wWZqpO69dnc9GUWcpGQLr7eVyKTs99WAPXR3kXpF4MVrHdBMEDfRfhytgomgAso96urN6eC8RaUpjX4Bf9HcAEJwfddZshin97XKJDmqCaqAfORNnf1e8hkfTIwYg1tvIpwemmEF4TkmOgK09N5dINyejyWMU8iWG8FqW5MfQ8A2DrtIdyGSKLH05s7H1dXyADjDECaC77QqLXTx7gWyHca3I0K92PuVFoOs5385vzqTYN3kVgFotSdXgoM8Zt5QIoj2lX4PYqm2TWsVp0s15JELikH8bNVIIMGiSSWJEWGU1PVEXD7V7cYepDb88korMjr3wbh6kZ76Q7F2RtfJqkd4hKw7B5OCX04b5eppkjL598iCpSUUx3X9C6fFavWj2DrHsv9DY86iCWBlcG08DRPKs9EPizCW4jNZtJcm3T7WlcI0MZMKOtsKOCWBZA0C9YnttNrp4eTsQ1U43StiIRPqp2K8rrQAu6etURH0RHedazHeeukTWI7iTG1dZpYk9EyittZ72qKXLNLhi5vJ9TlYw8O91vihB1nJwwA3B1WbiYhkqqRzoL0cQpXJMUsUlsoSP6Q70IMU92vEHbUmna5krESPLeJfQBKGQPNVVE63XYBh2TnvFTdi6koitu209wMFUnHZrzWj3UWGqsyTqqHbPl4RhRLFe24seRwV2SbUuLygBIdptKHnA3kutAbHzsWTT8UxOaiQzFV4auxounrgXj7MoMWEVKKS8AHkELPILGqFVFC8BZsfPC0WacSN5Rg5SaCvfs74hcsCQ3ghq9PyxEb2fbHUiaCjnsBcXqzQw9AjZJG4yX0ubEwicP0bKB6y3w4PUQqdouxH5y16OgkUjrZgodJfRLgP9vqGbHNDpj4yBuswluvCFBh38gBoSIQu11qtQmk43n4G8Dskn0DrJ32l2Gz35q5LaKT",
				"admin_password":   adminPassword,
				"nodes":            nodes,
			}
	}

	logger.Printf("%+v", nodes)
	doc, err := omClient.LoadDoc(planID, ctx)
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
