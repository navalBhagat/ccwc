package main

import (
	"bufio"
	"fmt"
	"log"
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

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	characterCount := 0
	lineCount := 0
	wordCount := 0
	reader := bufio.NewReader(file)

	for {
		contentByte, err := reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error reading file: %v", err)
		}
		characterCount++
		if contentByte == '\n' {
			lineCount++
		}
	}
	file.Seek(0, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCount++
	}

	return [4]int{characterCount, lineCount, wordCount, 0}
}
