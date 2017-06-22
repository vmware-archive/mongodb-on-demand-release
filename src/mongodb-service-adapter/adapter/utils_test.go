package adapter

import (
	"reflect"
	"testing"
)

func TestGenerateString(t *testing.T) {
	t.Parallel()

	s, err := GenerateString(8)
	if err != nil {
		t.Fatal(err)
	}

	if len(s) != 8 {
		t.Errorf("s, _ = GenerateString(%d); len(s) = %d, want %d", 8, len(s), 8)
	}
}

func TestNodesToCluster(t *testing.T) {
	t.Parallel()

	c, err := NodesToCluster([]string{
		"192.168.1.1", // Router
		"192.168.1.2", // Router

		"192.168.1.4", // Shard config
		"192.168.1.5", // Shard config

		"192.168.1.7", // Shard 0
		"192.168.1.8", // Shard 0
		"192.168.1.9", // Shard 0

		"192.168.1.10", // Shard 1
		"192.168.1.11", // Shard 1
		"192.168.1.12", // Shard 1

		"192.168.1.13", // Shard 2
		"192.168.1.14", // Shard 2
		"192.168.1.15", // Shard 2
	}, 2, 2, 3)

	if err != nil {
		t.Fatal(err)
	}

	want := &Cluster{
		Routers:       []string{"192.168.1.1", "192.168.1.2"},
		ConfigServers: []string{"192.168.1.4", "192.168.1.5"},
		Shards: [][]string{
			{"192.168.1.7", "192.168.1.8", "192.168.1.9"},
			{"192.168.1.10", "192.168.1.11", "192.168.1.12"},
			{"192.168.1.13", "192.168.1.14", "192.168.1.15"},
		},
	}

	if !reflect.DeepEqual(c, want) {
		t.Errorf("Cluster = %#v, want %#v", c, want)
	}
}
