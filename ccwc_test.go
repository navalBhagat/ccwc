package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	filename := "test.txt"
	expectedBytes := 342190
	expectedLines := 7145
	expectedWords := 58164
	expectedCharacters := 339292
	result := Count(filename)
	if result[0] != expectedBytes {
		t.Errorf("Expected bytes in %s to be %d but was %d", filename, expectedBytes, result[0])
	}
	if result[1] != expectedLines {
		t.Errorf("Expected lines in %s to be %d but was %d", filename, expectedLines, result[1])
	}
	if result[2] != expectedWords {
		t.Errorf("Expected words in %s to be %d but was %d", filename, expectedWords, result[2])
	}
	if result[3] != expectedCharacters {
		t.Errorf("Expected characters in %s to be %d but was %d", filename, expectedCharacters, result[3])
	}
}
