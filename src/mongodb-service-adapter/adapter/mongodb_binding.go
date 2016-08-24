package adapter

import (
	"fmt"
	"strings"

	"github.com/pivotal-cf/on-demand-service-broker-sdk/bosh"
	"github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"
	"gopkg.in/mgo.v2"
)

type Binder struct {
}

func (Binder) CreateBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest, requestParams serviceadapter.RequestParameters) (serviceadapter.Binding, error) {

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

  fmt.Printf(" **** Dial info *****")
	fmt.Printf("%s",dialInfo)
	fmt.Printf(" **** Dial info End *****")

	fmt.Printf(" **** Servers *****")
	fmt.Printf("%s",servers)
	fmt.Printf(" **** Servers End *****")
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


	fmt.Printf("This is the connection string :     mongodb://%s:%s@%s/%s?authSource=admin", username, password, strings.Join(servers, ","), username)

	return serviceadapter.Binding{
		Credentials: map[string]interface{}{
			"username": username,
			"password": password,
			"database": username,
			"servers":  servers,
			"uri":      fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin", username, password, strings.Join(servers, ","), username),
		},
	}, nil
}

func (Binder) DeleteBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest, requestParams serviceadapter.RequestParameters) error {

	// create an admin level user
	username := fmt.Sprintf("pcf_%v", encodeID(bindingID))

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

	fmt.Printf("%s",dialInfo)

	fmt.Printf("%s",servers)

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	adminDB := session.DB("admin")
	adminDB.RemoveUser(username)

	return nil
}
