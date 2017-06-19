package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

type DocContext struct {
	ID            string
	Key           string
	AdminPassword string
	Version       string
	Nodes         []string
	Shards        [][]string
	Password      string
}

func (oc *OMClient) LoadDoc(key string, ctx *DocContext) (string, error) {
	t, ok := plans[key]
	if !ok {
		return "", fmt.Errorf("plan %q not found", key)
	}

	if ctx.Password == "" {
		var err error
		ctx.Password, err = GenerateString(32)
		if err != nil {
			panic(err)
		}
	}

	b := bytes.Buffer{}
	if err := t.Execute(&b, ctx); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (oc *OMClient) CreateGroup(id string) (Group, error) {
	var group Group

	name := fmt.Sprintf("pcf_%s", id)
	resp, err := oc.doRequest("POST", "/api/public/v1.0/groups", strings.NewReader(`{
		"name": "`+name+`"
	}`))

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

func (oc *OMClient) GetGroup(GroupID string) (Group, error) {
	var group Group

	resp, err := oc.doRequest("GET", fmt.Sprintf("/api/public/v1.0/groups/%s", GroupID), nil)
	if err != nil {
		return group, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return group, err
	}

	if err = json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}

func (oc *OMClient) GetGroupHosts(GroupID string) (GroupHosts, error) {
	var groupHosts GroupHosts

	resp, err := oc.doRequest("GET", fmt.Sprintf("/api/public/v1.0/groups/%s/hosts", GroupID), nil)
	if err != nil {
		return groupHosts, err
	}
	defer resp.Body.Close()

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

func (oc *OMClient) ConfigureGroup(configurationDoc string, groupId string) error {
	url := fmt.Sprintf("/api/public/v1.0/groups/%s/automationConfig", groupId)
	body := strings.NewReader(configurationDoc)

	resp, err := oc.doRequest("PUT", url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(b))

	return nil
}

func (oc *OMClient) doRequest(method string, path string, body io.Reader) (*http.Response, error) {
	uri := fmt.Sprintf("%s%s", strings.TrimRight(oc.Url, "/"), path)
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(oc.Username, oc.ApiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("%s %s error: %v", method, uri, err)
		return nil, err
	}

	return res, nil
}
