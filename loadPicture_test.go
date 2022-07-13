package main

import (
	"testing"
)

func TestLoadPictures(t *testing.T) {
	pictures, err := LoadPictures("input")
	if err != nil {
		t.Error(err)
	}

	if len(pictures) != 2 {
		t.Error("Expected 2 pictures, got", len(pictures))
	}

	if pictures != ["abc.jpg", "test.jpg"] {
}
