package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(int, int) int) int {
	fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("10 + 2 =", compute(10, 2, add))
}
