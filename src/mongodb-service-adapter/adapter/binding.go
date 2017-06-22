package adapter

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
	"gopkg.in/mgo.v2"
)

type Binder struct {
	Logger *log.Logger
}

func (b *Binder) logf(msg string, v ...interface{}) {
	if b.Logger != nil {
		b.Logger.Printf(msg, v...)
	}
}

func (b Binder) CreateBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest, requestParams serviceadapter.RequestParameters) (serviceadapter.Binding, error) {

	// create an admin level user
	username := mkUsername(bindingID)
	password, err := GenerateString(32)
	if err != nil {
		return serviceadapter.Binding{}, err
	}

	properties := manifest.Properties["mongo_ops"].(map[interface{}]interface{})
	adminPassword := properties["admin_password"].(string)

	b.logf("properties: %v", properties)

	servers := make([]string, len(deploymentTopology["mongod_node"]))
	for i, node := range deploymentTopology["mongod_node"] {
		servers[i] = fmt.Sprintf("%s:28000", node)
	}

	routers := properties["routers"].(int)
	if routers != 0 {
		servers = servers[:routers]
	}

	dialInfo := &mgo.DialInfo{
		Addrs:     servers,
		Username:  "admin",
		Password:  adminPassword,
		Mechanism: "SCRAM-SHA-1",
		Database:  "admin",
		FailFast:  true,
	}

	b.logf("dialInfo: %v", dialInfo)

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return serviceadapter.Binding{}, err
	}
	defer session.Close()

	adminDB := session.DB("admin")

	// add user to admin database with admin privileges
	user := &mgo.User{
		Username: username,
		Password: password,
		Roles: []mgo.Role{
			mgo.RoleUserAdmin,
			mgo.RoleDBAdmin,
			mgo.RoleReadWrite,
		},
		OtherDBRoles: map[string][]mgo.Role{
			username: {
				mgo.RoleUserAdmin,
				mgo.RoleDBAdmin,
				mgo.RoleReadWrite,
			},
		},
	}
	adminDB.UpsertUser(user)

	url := fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin",
		username,
		password,
		strings.Join(servers, ","),
		username,
	)

	b.logf("url: %s", url)
	b.logf("username: %s", username)
	b.logf("password: %s", password)

	return serviceadapter.Binding{
		Credentials: map[string]interface{}{
			"username": username,
			"password": password,
			"database": username,
			"servers":  servers,
			"uri":      url,
		},
	}, nil
}

func (Binder) DeleteBinding(bindingID string, deploymentTopology bosh.BoshVMs, manifest bosh.BoshManifest, requestParams serviceadapter.RequestParameters) error {

	// create an admin level user
	username := mkUsername(bindingID)
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
		return err
	}
	defer session.Close()

	adminDB := session.DB("admin")
	adminDB.RemoveUser(username)

	return nil
}

func mkUsername(binddingID string) string {
	b64 := base64.StdEncoding.EncodeToString([]byte(binddingID))
	return fmt.Sprintf("pcf_%x", md5.Sum([]byte(b64)))
}
