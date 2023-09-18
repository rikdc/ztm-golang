//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	count int
	sync.Mutex
}

func getWords(line string) []string {
	return strings.Split(line, " ")
}

func countLeters(word string) int {
	letters := 0

	for _, letter := range word {
		if unicode.IsLetter(letter) {
			letters += 1
		}
	}

	return letters
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	totalLetters := Count{}
	var wg sync.WaitGroup

	for scanner.Scan() {

		go func() {
			words := getWords(scanner.Text())

			for _, word := range words {
				wordCopy := word

				wg.Add(1)

				go func() {
					totalLetters.Lock()
					defer totalLetters.Unlock()
					defer wg.Done()

					totalLetters.count += countLeters(wordCopy)
				}()
			}
		}()
	}

	fmt.Println("waiting for goroutines to finish")
	wg.Wait()

	totalLetters.Lock()
	totalLettersResult := totalLetters.count
	defer totalLetters.Unlock()
	fmt.Printf("Total letters: %d\n", totalLettersResult)

	fmt.Println("done!")
}
