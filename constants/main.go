package main

import "fmt"

const name = "John Doe"

func main() {
	const age = 30

	const (
		country = "USA"
		city    = "New York"
	)

	// Error: cannot assign to name
	// name = "John Doe"
	// Error: cannot assign to age
	// age = 31

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)
	fmt.Println("City:", city)
}
