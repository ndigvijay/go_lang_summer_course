package main

import "fmt"

// Define a function type for the rectangle function
// type rectangle func(int, int) int

// Define the Shape struct
type Shape struct {
	length  int
	breadth int
	color   string
	rect    rectangle
}

// Implement the rectangle function
func rectangle(length int, breadth int) int {
	return length * breadth
}

func main() {
	// Create a Shape instance
	shape := Shape{length: 10, breadth: 20, color: "blue", rect: rectangle}

	// Calculate the area using the rect function
	area := shape.rect(shape.length, shape.breadth)

	// Print the result
	fmt.Printf("The area of the %s rectangle is %d\n", shape.color, area)
}
