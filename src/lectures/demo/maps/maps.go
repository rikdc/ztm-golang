package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["bread"] = 1
	shoppingList["milk"] += 1
	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println(shoppingList)

	fmt.Println("We need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Cereal:", cereal)
	}

	var totalItems = 0
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("Total items:", totalItems)
}
