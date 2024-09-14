package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	args := os.Args[1:]

	counts, filename, flag := GetCounts(args)
	switch flag {
	case "-c":
		fmt.Println(counts[0], filename)
	case "-l":
		fmt.Println(counts[1], filename)
	case "-w":
		fmt.Println(counts[2], filename)
	case "-m":
		fmt.Println(counts[3], filename)
	default:
		fmt.Println(counts[1], counts[2], counts[0], filename)
	}
}

func GetCounts(args []string) ([4]int, string, string) {
	var flag string
	var counts [4]int
	var filename string

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		if len(args) == 1 {
			flag = args[0]
		} else if len(args) > 1 {
			fmt.Println("Usage: cat <filename> | ccwc [-c|-l|-w|-m]")
			os.Exit(1)
		}
		counts = CountFromStdIn()
		filename = ""
	} else {
		if len(args) == 1 {
			filename = args[0]
		} else if len(args) == 2 {
			flag = args[0]
			filename = args[1]
		} else {
			fmt.Println("Usage: ccwc [-c|-l|-w|-m] <filename>")
			os.Exit(1)
		}
		counts = CountFromFile(filename)
	}

	return counts, filename, flag
}

func CountFromFile(filename string) [4]int {

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

func CountFromStdIn() [4]int {
	byteCount := 0
	lineCount := 0
	wordCount := 0
	characterCount := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	inWord := false
	for scanner.Scan() {
		r := scanner.Text()
		byteCount += len(r)
		if r == "\n" {
			lineCount++
		}
		if unicode.IsSpace([]rune(r)[0]) {
			if inWord {
				wordCount++
				inWord = false
			}
		} else {
			inWord = true
		}
		characterCount++
	}
	if inWord {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}

	return [4]int{byteCount, lineCount, wordCount, characterCount}
}
