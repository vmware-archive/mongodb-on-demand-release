package adapter

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/AsGz/httpAuthClient"
	"github.com/aymerick/raymond"
	"github.com/nu7hatch/gouuid"
)

type OMClient struct {
	Url      string
	Username string
	ApiKey   string
}

type Group struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	AgentAPIKey string         `json:"agentApiKey"`
	HostCounts  map[string]int `json:"hostCounts"`
}

type GroupHosts struct {
	TotalCount int `json:"totalCount"`
}

func (oc OMClient) LoadDoc(key string, ctx map[string]interface{}) (string, error) {
	raymond.RegisterHelper("password", func() string {
		return oc.RandomString(32)
	})

	raymond.RegisterHelper("isConfig", func(index int) bool {
		return index >= 9 && index < 12
	})

	raymond.RegisterHelper("isInShard", func(index int) bool {
		return index < 12
	})

	raymond.RegisterHelper("hasStorage", func(index int) bool {
		return index < 12
	})

	raymond.RegisterHelper("processType", func(index int) string {
		if index > 11 && index < 15 {
			return "mongos"
		} else {
			return "mongod"
		}
	})

	raymond.RegisterHelper("hasShardedCluster", func(index int) bool {
		return index > 11 && index < 15
	})

	raymond.RegisterHelper("div", func(val int, div int) int {
		return val / div
	})

	tpl, ok := plans[key]
	if !ok {
		return "", fmt.Errorf("plan %q not found", key)
	}

	result, err := raymond.Render(tpl, ctx)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (oc OMClient) CreateGroup() (Group, error) {
	var group Group

	u, err := uuid.NewV4()
	if err != nil {
		return group, err
	}

	groupName := fmt.Sprintf("pcf_%s", u)
	body := strings.NewReader(fmt.Sprintf("{\"name\": \"%s\"}", groupName))

	resp, err := oc.doRequest("POST", "/api/public/v1.0/groups", body)
	if err != nil {
		return group, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return group, err
	}

	if err = json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}

func (oc OMClient) GetGroup(GroupID string) (Group, error) {
	var group Group

	resp, err := oc.doRequest("GET", fmt.Sprintf("/api/public/v1.0/groups/%s", GroupID), nil)
	if err != nil {
		return group, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return group, err
	}

	if err = json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}

func (oc OMClient) GetGroupHosts(GroupID string) (GroupHosts, error) {
	var groupHosts GroupHosts

	resp, err := oc.doRequest("GET", fmt.Sprintf("/api/public/v1.0/groups/%s/hosts", GroupID), nil)
	if err != nil {
		return groupHosts, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return groupHosts, err
	}

	err = json.Unmarshal(b, &groupHosts)
	if err != nil {
		return groupHosts, err
	}

	return groupHosts, nil
}

func (oc OMClient) ConfigureGroup(configurationDoc string, groupId string) error {
	url := fmt.Sprintf("/api/public/v1.0/groups/%s/automationConfig", groupId)
	body := strings.NewReader(configurationDoc)

	resp, err := oc.doRequest("PUT", url, body)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(b))

	return nil
}

func (oc OMClient) doRequest(method string, path string, body io.Reader) (*http.Response, error) {
	uri := fmt.Sprintf("%s%s", strings.TrimRight(oc.Url, "/"), path)
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	err = httpAuthClient.ApplyHttpDigestAuth(oc.Username, oc.ApiKey, uri, req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("%s %s error: %v", method, uri, err)
		return nil, err
	}

	return res, nil
}

func (oc OMClient) RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// func (oc OMClient) PostDoc(url string, username string, apiKey string) {
//
// }