package adapter

import "testing"

func TestOMClient_LoadDoc(t *testing.T) {
	t.Parallel()

	c := &OMClient{}
	for p, ctx := range map[Plan]*DocContext{
		PlanStandalone: {
			ID:            "d3a98gf1",
			Key:           "key",
			AdminPassword: "pwd",
			Nodes:         []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"},
			Version:       "3.2.11",
		},

		PlanShardedSet: {
			ID:            "d3a98gf1",
			Key:           "key",
			AdminPassword: "pwd",
			Nodes:         []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"},
			Version:       "3.2.11",
			Shards:        [][]string{{"192.168.1.1", "192.168.1.2"}, {"192.168.1.3", "192.168.1.4"}},
		},

		PlanSingleReplicaSet: {
			ID:            "d3a98gf1",
			Key:           "key",
			AdminPassword: "pwd",
			Nodes:         []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"},
			Version:       "3.2.11",
		},
	} {
		t.Run(string(p), func(t *testing.T) {
			_, err := c.LoadDoc(p, ctx)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
