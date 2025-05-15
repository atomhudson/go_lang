package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop
	i := 1
	for i <= 10 {
		fmt.Printf("%d  ", i)
		i++
	}
	// for loop
	for j := 1; j <= 10; j++ {
		fmt.Printf("%d  ", j)
	}
	// infinite loop
	for {
		println("Infinite Loop")
		break
	}
	// range loop
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	// break and continue
	for k := 1; k <= 10; k++ {
		if k == 5 {
			continue // skip the rest of the loop when k is 5
		}
		if k == 8 {
			break // exit the loop when k is 8
		}
		fmt.Printf("%d ", k)
	}
	// nested loop
outerLoop:
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			if a == 2 && b == 2 {
				break outerLoop // exit the outer loop when a is 2 and b is 2
			}
			fmt.Printf("Outer: %d, Inner: %d\n", a, b)
		}
	}
}
