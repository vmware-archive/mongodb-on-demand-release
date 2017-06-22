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

func TestSortAddresses(t *testing.T) {
	t.Parallel()

	s1 := []string{"192.168.1.02", "192.168.1.1", "192.168.1.10"}
	s2 := []string{"192.168.1.1", "192.168.1.02", "192.168.1.10"}

	SortAddresses(s1)
	if !reflect.DeepEqual(s1, s2) {
		t.Errorf("SortAddresses = %v, want %v", s1, s2)
	}
}
