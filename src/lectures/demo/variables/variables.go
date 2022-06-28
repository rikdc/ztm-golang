package main

import "fmt"

func main() {
	var myName = "Richard"
	fmt.Println("My name is", myName)

	var name string = "Kathy"
	fmt.Println("name =", name)

	username := "admin"
	fmt.Println("username =", username)

	var sum int
	fmt.Println("The sum is", sum)

	part1, other := 1, 5
	fmt.Println("part1 is", part1, "other is", other)

	part2, other := 2, 0
	fmt.Println("part2 is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum =", sum)

	var (
		lessonName = "variables"
		lessonType = "demo"
	)

	fmt.Println("lessonName =", lessonName, "lessonType =", lessonType)

}
