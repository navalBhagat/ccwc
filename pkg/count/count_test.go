package count

import (
	"bufio"
	"os"
	"testing"
)

func TestCountFromFile(t *testing.T) {
	filename := "../../testdata/test.txt"
	expected := [4]int{342190, 7145, 58164, 339292}

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := CountWithScanner(scanner)

	checkCounts(t, filename, result, expected)
}

func TestCountFromStdIn(t *testing.T) {
	expected := [4]int{342190, 7145, 58164, 339292}

	// Simulate input from stdin by opening the test file and redirecting stdin
	file, err := os.Open("../../testdata/test.txt")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	origStdin := os.Stdin
	defer func() { os.Stdin = origStdin }()
	os.Stdin = file

	scanner := bufio.NewScanner(os.Stdin)
	result := CountWithScanner(scanner)

	checkCounts(t, "stdin", result, expected)
}

// Helper function to check counts and avoid redundancy
func checkCounts(t *testing.T, source string, result [4]int, expected [4]int) {
	if result[0] != expected[0] {
		t.Errorf("Expected bytes in %s to be %d but got %d", source, expected[0], result[0])
	}
	if result[1] != expected[1] {
		t.Errorf("Expected lines in %s to be %d but got %d", source, expected[1], result[1])
	}
	if result[2] != expected[2] {
		t.Errorf("Expected words in %s to be %d but got %d", source, expected[2], result[2])
	}
	if result[3] != expected[3] {
		t.Errorf("Expected characters in %s to be %d but got %d", source, expected[3], result[3])
	}
}
