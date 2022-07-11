package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {

	c := Counter{}

	hello := "Hello"
	world := "world"

	fmt.Println(hello, world)

	replace(&hello, "Hi", &c)

	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &c)
	fmt.Println(phrase)

}
