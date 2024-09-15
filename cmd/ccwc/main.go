package main

import (
	"bufio"
	"ccwc/pkg/count"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	counts, filename, flag := GetCounts(args)
	outputCounts(counts, filename, flag)
}

func GetCounts(args []string) ([4]int, string, string) {
	flag, filename := parseArgs(args)
	var counts [4]int

	if isInputFromPipe() {
		counts = count.CountWithScanner(bufio.NewScanner(os.Stdin))
	} else {
		counts = CountFromFile(filename)
	}

	return counts, filename, flag
}

func parseArgs(args []string) (string, string) {
	var flag, filename string

	switch len(args) {
	case 1:
		if isInputFromPipe() {
			flag = args[0]
		} else {
			filename = args[0]
		}
	case 2:
		flag, filename = args[0], args[1]
	default:
		printUsageAndExit()
	}

	return flag, filename
}

func isInputFromPipe() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func printUsageAndExit() {
	fmt.Println("Usage: ccwc [-c|-l|-w|-m] <filename> or cat <filename> | ccwc [-c|-l|-w|-m]")
	os.Exit(1)
}

func CountFromFile(filename string) [4]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	return count.CountWithScanner(bufio.NewScanner(file))
}

func outputCounts(counts [4]int, filename, flag string) {
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
