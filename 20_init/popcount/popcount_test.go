package popcount

import "testing"

func TestPopCount(t *testing.T) {
	want := 99
	if got := PopCount(99); got != want {
		t.Errorf("PopCount() = %q, want %q", got, want)
	}
}
