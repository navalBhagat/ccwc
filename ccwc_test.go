package main

import (
	"testing"
)

func TestCharacterCount(t *testing.T) {
	filename := "test.txt"
	expected := 342190
	result := Count(filename)
	if result[0] != expected {
		t.Errorf("Expected characters in %s to be %d but was %d", filename, expected, result[0])
	}
}
