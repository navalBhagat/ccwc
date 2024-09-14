package main

import (
	"os"
	"testing"
)

func TestCountFromFile(t *testing.T) {
	filename := "test.txt"
	expectedBytes := 342190
	expectedLines := 7145
	expectedWords := 58164
	expectedCharacters := 339292
	result := CountFromFile(filename)
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

func TestCountFromStdIn(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	origStdin := os.Stdin
	defer func() { os.Stdin = origStdin }()

	// Redirect stdin to the test file
	os.Stdin = file

	result := CountFromStdIn()

	expectedBytes := 342190
	expectedLines := 7145
	expectedWords := 58164
	expectedCharacters := 339292
	if result[0] != expectedBytes {
		t.Errorf("Expected bytes in %s to be %d but was %d", "", expectedBytes, result[0])
	}
	if result[1] != expectedLines {
		t.Errorf("Expected lines in %s to be %d but was %d", "", expectedLines, result[1])
	}
	if result[2] != expectedWords {
		t.Errorf("Expected words in %s to be %d but was %d", "", expectedWords, result[2])
	}
	if result[3] != expectedCharacters {
		t.Errorf("Expected characters in %s to be %d but was %d", "", expectedCharacters, result[3])
	}
}
