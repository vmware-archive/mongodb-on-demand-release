package main

import (
	"flag"

	"github.com/cf-platform-eng/mongodb-on-demand-release/src/mongodb-config-agent/agent"
)

func main() {
	configAgent := agent.ConfigAgent{}

	url := flag.String("url", "", "MOM URL")
	username := flag.String("username", "", "MOM Username")
	apiKey := flag.String("api-key", "", "MOM API Key")
	groupID := flag.String("group", "", "MOM Group ID")
	planID := flag.String("plan", "", "The name of the service plan")
	nodeAddresses := flag.String("nodes", "", "Comma separated list of addresses")
	adminPassword := flag.String("admin-password", "", "Admin password for the mongo instance")
	engineVersion := flag.String("engine-version", "", "Engine version")
	replicas := flag.Int("replicas", 0, "replicas per shard")

	flag.Parse()

	configAgent.PollAndConfigureGroup(
		*url,
		*username,
		*apiKey,
		*groupID,
		*planID,
		*nodeAddresses,
		*adminPassword,
		*engineVersion,
		*replicas,
	)
}
