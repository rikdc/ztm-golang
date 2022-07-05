//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a Coordinate
	b Coordinate
}

func width(rect Rectangle) int {
	return rect.b.x - rect.a.x
}

func length(rect Rectangle) int {
	return rect.a.y - rect.b.y
}

func area(rect Rectangle) int {
	fmt.Println("Length is", length(rect))
	fmt.Println("Width is", width(rect))
	return length(rect) * width(rect)
}

func perimeter(rect Rectangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

func describeRect(rect Rectangle) {
	fmt.Println("The area is", area(rect))
	fmt.Println("The permiter is", perimeter(rect))
}

func main() {

	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	describeRect(rect)

	rect.a.y *= 2
	rect.b.x *= 2
	describeRect(rect)
}
