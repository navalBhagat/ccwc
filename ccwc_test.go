package main

import (
	"testing"
)

func TestCharacterCount(t *testing.T) {
	filename := "test.txt"
	expected := 342190
	result := Count(filename)
	if result != expected {
		t.Errorf("Count(%s) = %d; want %d", filename, result, expected)
	}
}
