package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	args := os.Args[1:]

	var flag, filename string
	if len(args) == 1 {
		filename = args[0]
	} else if len(args) == 2 {
		flag = args[0]
		filename = args[1]
	} else {
		fmt.Println("Usage: ccwc [-c|-l|-w|-m] <filename>")
		os.Exit(1)
	}

	counts := Count(filename)
	switch flag {
	case "-c":
		fmt.Println(counts[0], filename)
	case "-l":
		fmt.Println(counts[1], filename)
	case "-w":
		fmt.Println(counts[2], filename)
	case "-m":
		fmt.Println(counts, filename)
	default:
		fmt.Println(counts[1], counts[2], counts[0], filename)
	}
}

func Count(filename string) [4]int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	byteCount := 0
	lineCount := 0
	wordCount := 0
	characterCount := 0

	reader := bufio.NewReader(file)
	for {
		contentByte, err := reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error reading file: %v", err)
		}
		byteCount++
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

	file.Seek(0, 0)
	reader = bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error reading file: %v", err)
		}
		for _, r := range line {
			if utf8.ValidRune(r) {
				characterCount++
			}
		}
	}

	return [4]int{byteCount, lineCount, wordCount, characterCount}
}
