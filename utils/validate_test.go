package main

import "testing"

func TestListIsEqual(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	if !listIsEqual(a, b) {
		t.Error("Expected", a, "to equal", b)
	}
}
