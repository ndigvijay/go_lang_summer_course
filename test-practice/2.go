package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{Name: "Alice", Age: 30} // Initialize p1 with Name "Alice" and Age 30
	p2 := p1                             // Assign p1 to p2 (shallow copy)
	p2.Age = 25                          // Modify Age of p2

	fmt.Println("Person 1:", p1) // Print p1
	fmt.Println("Person 2:", p2) // Print p2
}
