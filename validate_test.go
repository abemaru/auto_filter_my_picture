package main

import "testing"

func TestListIsEqual_正常系(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "c"}
	if !listIsEqual(a, b) {
		t.Error("Expected", a, "to equal", b)
	}
}

func TestListIsEqual_異常系(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b", "d"}
	if listIsEqual(a, b) {
		t.Error("Expected", a, "to not equal", b)
	}
}
