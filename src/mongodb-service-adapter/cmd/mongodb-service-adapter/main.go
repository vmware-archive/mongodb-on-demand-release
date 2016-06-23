package main

import (
	"log"
	"os"

	"github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"
	"mongodb-service-adapter/adapter"
)

func main() {
	logger := log.New(os.Stderr, "[mongodb-service-adapter] ", log.LstdFlags)
	//sa := adapter.Adapter{}
	manifestGenerator := adapter.ManifestGenerator{}
	binder := adapter.Binder{}
	serviceadapter.HandleCommandLineInvocation(os.Args, manifestGenerator, binder, nil, logger)
}
