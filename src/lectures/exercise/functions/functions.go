//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

func greeting(name string) {
	fmt.Println("Hello,", name)
}

func returnAMessage() string {
	return "I am a message!"
}

func add(a, b, c int) int {
	return a + b + c
}

func anyNumber() int {
	r := rand.New(rand.NewSource(99))
	return r.Intn(5000)
}

func anyTwoNumbers() (int, int) {
	r := rand.New(rand.NewSource(99))
	return r.Intn(5000), r.Intn(5000)
}

func main() {
	greeting("Bob")
	fmt.Println(returnAMessage())

	theSum := add(1, 2, 3)
	fmt.Println(theSum)

	a := anyNumber()
	b, c := anyTwoNumbers()

	fmt.Println("The sum of all three numbers:", add(a, b, c))

}
