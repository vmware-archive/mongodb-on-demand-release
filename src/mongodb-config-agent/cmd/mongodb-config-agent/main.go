package main

import (
	"flag"

	"github.com/cf-platform-eng/mongodb-on-demand-release/src/mongodb-config-agent/config"
)

func main() {
	configAgent := config.ConfigAgent{}

	url := flag.String("url", "", "MOM URL")
	username := flag.String("username", "", "MOM Username")
	apiKey := flag.String("api-key", "", "MOM API Key")
	groupID := flag.String("group", "", "MOM Group ID")
	planID := flag.String("plan", "", "The name of the service plan")
	nodeAddresses := flag.String("nodes", "", "Comma separated list of addresses")

	flag.Parse()

	configAgent.PollAndConfigureGroup(
		*url,
		*username,
		*apiKey,
		*groupID,
		*planID,
		*nodeAddresses,
	)
}
