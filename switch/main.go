package main

import "fmt"

func main() {
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		case float64:
			fmt.Println("it's a float")
		default:
			fmt.Println("unknown type: ", t)
		}
	}
	whoAmI(1)
	whoAmI("hello")
	whoAmI(true)
	whoAmI(3.14)
	whoAmI([]int{1, 2, 3})
}
