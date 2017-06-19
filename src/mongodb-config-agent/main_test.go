package main

import (
	"reflect"
	"testing"
)

func TestNodesToShards(t *testing.T) {
	t.Parallel()

	nodes := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"}
	want := [][]string{{"192.168.1.1", "192.168.1.2"}, {"192.168.1.3", "192.168.1.4"}}

	got, err := nodesToShards(nodes, 2)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("nodesToShards(%v, %d) = %v, want %v", nodes, 2, got, want)
	}
}
