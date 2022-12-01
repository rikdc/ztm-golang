package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element found at position %d", index))
	}

	return s.values[index], nil
}

func main() {
	s := Stuff{}

	value, err := s.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
