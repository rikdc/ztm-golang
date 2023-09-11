//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(string)

func iterateText(lines []string, fn LineCallback) {
	for _, line := range lines {
		fn(line)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	//  - Number of letters
	var letterCount, digitCount, spaceCount, punctuationCount int = 0, 0, 0, 0

	iterateText(lines, func(line string) {
		for _, r := range line {
			if unicode.IsLetter(r) {
				letterCount++
			}
			if unicode.IsDigit(r) {
				digitCount++
			}
			if unicode.IsSpace(r) {
				spaceCount++
			}
			if unicode.IsPunct(r) {
				punctuationCount++
			}
		}
	})

	fmt.Printf("Number of letters: %v\n", letterCount)
	fmt.Printf("Number of digits: %v\n", digitCount)
	fmt.Printf("Number of spaces: %v\n", spaceCount)
	fmt.Printf("Number of punctuation marks: %v\n", punctuationCount)
}
