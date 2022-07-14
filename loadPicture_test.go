package main

import (
	"testing"
)

func TestLoadPictures_正常系(t *testing.T) {
	pictures, err := LoadPictures("input")
	if err != nil {
		t.Error(err)
	}
	expected := []string{"input/a.jpg", "input/b.jpg"}

	if !listIsEqual(pictures, expected) {
		t.Error("Expected", pictures, "to equal", expected)
	}
}
