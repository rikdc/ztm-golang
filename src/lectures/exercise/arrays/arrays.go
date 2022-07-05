//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func totals(items [4]Product) (int, float64) {
	var total = 0.0
	var count = 0

	for i := 0; i < len(items); i++ {

		if items[i].name != "" {
			total += items[i].price
			count += 1
		}
	}

	return count, total
}

func lastItemName(items [4]Product) string {
	return items[len(items)-1].name
}

func printStats(items [4]Product) {

	count, total := totals(items)

	fmt.Println("List total is", total)
	fmt.Println("Item Count is", count)
	fmt.Println("Last item", lastItemName(items))
}

func main() {

	items := [...]Product{
		{name: "Pencil", price: 2.99},
		{name: "Paper", price: 5.34},
		{name: "Pen", price: 14.23},
		{name: "Ink", price: 1.32},
	}

	printStats(items)

	items[3] = Product{name: "Goat", price: 2.99}

	printStats(items)
}
