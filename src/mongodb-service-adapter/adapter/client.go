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

	"github.com/cf-platform-eng/mongodb-on-demand-release/src/mongodb-service-adapter/digest"
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
	ID                   string
	Key                  string
	AdminPassword        string
	Version              string
	CompatibilityVersion string
	Nodes                []string
	Cluster              *Cluster
	Password             string
}

type Cluster struct {
	Routers       []string
	ConfigServers []string
	Shards        [][]string
}

func (oc *OMClient) LoadDoc(p string, ctx *DocContext) (string, error) {
	t, ok := plans[p]
	if !ok {
		return "", fmt.Errorf("plan %q not found", p)
	}

	if ctx.Password == "" {
		var err error
		ctx.Password, err = GenerateString(32)
		if err != nil {
			panic(err)
		}
	}

	if strings.HasPrefix(ctx.Version, "3.4") {
		ctx.CompatibilityVersion = "3.4"
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
	b, err := oc.doRequest("POST", "/api/public/v1.0/groups", strings.NewReader(`{
		"name": "`+name+`"
	}`))
	if err != nil {
		return group, err
	}

	if err = json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}

func (oc *OMClient) GetGroup(groupID string) (Group, error) {
	var group Group

	b, err := oc.doRequest("GET", fmt.Sprintf("/api/public/v1.0/groups/%s", groupID), nil)
	if err != nil {
		return group, err
	}

	if err = json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}

func (oc *OMClient) GetGroupHosts(groupID string) (GroupHosts, error) {
	var groupHosts GroupHosts

	b, err := oc.doRequest("GET", fmt.Sprintf("/api/public/v1.0/groups/%s/hosts", groupID), nil)
	if err != nil {
		return groupHosts, err
	}

	if err = json.Unmarshal(b, &groupHosts); err != nil {
		return groupHosts, err
	}
	return groupHosts, nil
}

func (oc *OMClient) ConfigureGroup(configurationDoc string, groupID string) error {
	u := fmt.Sprintf("/api/public/v1.0/groups/%s/automationConfig", groupID)
	b, err := oc.doRequest("PUT", u, strings.NewReader(configurationDoc))
	if err != nil {
		return err
	}
	log.Println(string(b))

	return nil
}

func (oc *OMClient) doRequest(method string, path string, body io.Reader) ([]byte, error) {
	uri := fmt.Sprintf("%s%s", oc.Url, path)
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if err = digest.ApplyDigestAuth(oc.Username, oc.ApiKey, uri, req); err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("%s %s error: %v", method, uri, err)
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 && res.StatusCode >= 300 {
		return nil, fmt.Errorf("%s %s request error: code=%d body=%q", method, path, res.StatusCode, b)
	}
	return b, nil
}
