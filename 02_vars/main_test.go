package main

import "testing"

func TestFunctions(t *testing.T) {
	s := "stringst"
	suf := "st"
	wanthasPrefix := true
	if got := HasSuffix(s, suf); got != wanthasPrefix {
		t.Errorf("HasSuffix() = %v, want %v", got, wanthasPrefix)
	}

}
