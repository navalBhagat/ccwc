package count

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func CountWithScanner(scanner *bufio.Scanner) [4]int {
	byteCount, lineCount, wordCount, characterCount := 0, 0, 0, 0
	inWord := false

	scanner.Split(bufio.ScanRunes)

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
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	return [4]int{byteCount, lineCount, wordCount, characterCount}
}
