package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	filename := "test.txt"
	expectedCharacters := 342190
	expectedLines := 7145
	expectedWords := 58164
	result := Count(filename)
	if result[0] != expectedCharacters {
		t.Errorf("Expected characters in %s to be %d but was %d", filename, expectedCharacters, result[0])
	}
	if result[1] != expectedLines {
		t.Errorf("Expected lines in %s to be %d but was %d", filename, expectedLines, result[0])
	}
	if result[2] != expectedWords {
		t.Errorf("Expected lines in %s to be %d but was %d", filename, expectedWords, result[0])
	}
}
