package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Count("test.txt"), "test.txt")
}

func Count(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	characterCount := len(file)
	return characterCount
}
