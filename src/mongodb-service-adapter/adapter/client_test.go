package adapter

import (
	"fmt"
	"testing"
)

func TestOMClient_LoadDoc(t *testing.T) {
	t.Parallel()

	ctx := &DocContext{
		ID:            "d3a98gf1",
		Key:           "key",
		AdminPassword: "pwd",
		Nodes:         []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"},
		Version:       "3.2.11",
		Shards:        [][]string{{"192.168.1.1", "192.168.1.2"}, {"192.168.1.3", "192.168.1.4"}},
	}

	c := &OMClient{}
	d, err := c.LoadDoc(PlanShardedSet, ctx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(d)
}
