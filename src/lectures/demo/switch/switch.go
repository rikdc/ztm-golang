package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {
	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("Moderately priced item")
	default:
		fmt.Println("Expensive item")
	}

	ticket := FirstClass
	switch ticket {
	case Economy:
		fmt.Println("Economy Ticket")
	case Business:
		fmt.Println("Business Ticket")
	case FirstClass:
		fmt.Println("FirstClass Ticket")
	default:
		fmt.Println("Other Seating")

	}
}
