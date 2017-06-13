package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cf-platform-eng/mongodb-on-demand-release/src/mongodb-service-adapter/adapter"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

func main() {
	if len(os.Args) < 2 {
		usage("<generate-manifest|create-binding|delete-binding|dashboard-url>")
	}

	switch os.Args[1] {
	case "generate-manifest":
		if len(os.Args) != 7 {
			usage("generate-manifest serviceDeploymentJSON planJSON argsJSON previousManifestYAML previousPlanJSON")
		}
	case "create-binding":
		if len(os.Args) != 6 {
			usage("create-binding bindingID boshVMsJSON manifestYAML requestParams")
		}
	case "delete-binding":
		if len(os.Args) != 6 {
			usage("delete-binding bindingID boshVMsJSON manifestYAML unbindingRequestParams")
		}
	case "dashboard-url":
		if len(os.Args) != 5 {
			usage("dashboard-url instanceID planJSON manifestYAML")
		}
	default:
		usage("<generate-manifest|create-binding|delete-binding|dashboard-url>")
	}

	logger := log.New(os.Stderr, "[mongodb-config-agent] ", log.LstdFlags)

	serviceadapter.HandleCommandLineInvocation(
		os.Args,
		&adapter.ManifestGenerator{logger},
		&adapter.Binder{logger},
		&adapter.DashboardURLGenerator{},
	)
}

func usage(msg string) {
	exe, err := os.Executable()
	if err != nil {
		exe = os.Args[0]
	}

	fmt.Fprintf(os.Stderr, "usage: %s %s\n", exe, msg)
	os.Exit(1)
}
