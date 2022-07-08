//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(parts []Part) {
	fmt.Println("---")
	for i := 0; i < len(parts); i++ {
		part := parts[i]
		fmt.Println(part)
	}
}

func main() {

	// Create an assembly line having any three parts
	assemblyLine := make([]Part, 3)
	assemblyLine[0] = "Part 1"
	assemblyLine[1] = "Part 2"
	assemblyLine[2] = "Part 3"

	printAssemblyLine(assemblyLine)

	// Add two new parts to the line
	assemblyLine = append(assemblyLine, "Part 4", "Part 5")

	printAssemblyLine(assemblyLine)

	// Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:5]
	printAssemblyLine(assemblyLine)

}
