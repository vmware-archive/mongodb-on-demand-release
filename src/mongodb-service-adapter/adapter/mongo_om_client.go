package adapter

import (
	"encoding/json"
	"strings"
  "log"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/AsGz/httpAuthClient"
	"net/http"
	"io/ioutil"
	"io"
)

type OMClient struct {
	Url 			string
	Username 	string
	ApiKey 		string
}

type Group struct {
  ID          string `json:"id"`
  Name        string `json:"name"`
  AgentAPIKey string `json:"agentApiKey"`
}

func (oc OMClient) LoadDoc(key string) string {
  docs := map[string]string {
    "single_node": "om_cluster_docs/3_2_cluster.json",
    "single_replica_set": "om_cluster_docs/replica-set.json",
    "sharded_cluster": "om_cluster_docs/3_2_cluster.json",
  }

  path := docs[key]
  asset, _ := Asset(path)
  return string(asset)
}

func (oc OMClient) CreateGroup() (Group, error) {

	u, err := uuid.NewV4()
	groupName := fmt.Sprintf("pcf_%s", u)
  body := strings.NewReader(fmt.Sprintf("{\"name\": \"%s\"}", groupName))

  var group Group

	resp, err := oc.doRequest("POST", "/api/public/v1.0/groups", body)

  if err != nil {
    log.Fatalf("could not post: %v", err)
		return group, err
  }

	var b []byte
  b, err = ioutil.ReadAll(resp.Body)
  err = json.Unmarshal(b, &group)

  return group, nil
}

func (oc OMClient) doRequest(method string, path string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", oc.Url, path), body)
	req.Header.Set("Content-Type", "application/json")

	err = httpAuthClient.ApplyHttpDigestAuth(oc.Username, oc.ApiKey, fmt.Sprintf("%s%s", oc.Url, path), req)

	if err != nil {
    log.Fatal(err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

  if err != nil {
    log.Fatalf("could not post: %v", err)
		return nil, err
  }

	return resp, nil
}

// func (oc OMClient) PostDoc(url string, username string, apiKey string) {
//
// }
