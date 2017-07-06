package adapter

import (
	"encoding/json"
	"testing"
)

func TestOMClient_LoadDoc(t *testing.T) {
	t.Parallel()

	c := &OMClient{}
	for p, ctx := range map[string]*DocContext{
		PlanStandalone: {
			ID:            "d3a98gf1",
			Key:           "key",
			AdminPassword: "pwd",
			Nodes:         []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"},
			Version:       "3.2.11",
		},

		PlanShardedCluster: {
			ID:            "d3a98gf1",
			Key:           "key",
			AdminPassword: "pwd",
			Version:       "3.4.1",
			Cluster: &Cluster{
				Routers:       []string{"192.168.1.1", "192.168.1.2"},
				ConfigServers: []string{"192.168.1.3", "192.168.1.4"},
				Shards: [][]string{
					{"192.168.0.10", "192.168.0.11", "192.168.0.12"},
					{"192.168.0.13", "192.168.0.14", "192.168.0.15"},
				},
			},
		},

		PlanReplicaSet: {
			ID:            "d3a98gf1",
			Key:           "key",
			AdminPassword: "pwd",
			Nodes:         []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"},
			Version:       "3.2.11",
		},
	} {
		t.Run(string(p), func(t *testing.T) {
			s, err := c.LoadDoc(p, ctx)
			if err != nil {
				t.Fatal(err)
			}

			// validate json output
			if err := json.Unmarshal([]byte(s), &map[string]interface{}{}); err != nil {
				t.Fatal(err)
			}
		})
	}
}
