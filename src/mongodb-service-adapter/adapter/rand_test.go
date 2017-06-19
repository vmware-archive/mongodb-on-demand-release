package adapter

import "testing"

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
