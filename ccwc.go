package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	flag := args[0]
	filename := args[1]
	switch flag {
	case "-c":
		fmt.Println(Count(filename)[0], "test.txt")
	case "-l":
		fmt.Println(Count(filename)[1], "test.txt")
	case "-w":
		fmt.Println(Count(filename)[2], "test.txt")
	case "-m":
		fmt.Println(Count(filename)[3], "test.txt")
	default:
		fmt.Println(Count(flag), "test.txt")
	}
}

func Count(filename string) [4]int {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	characterCount := len(file)
	return [4]int{characterCount, 0, 0, 0}
}
