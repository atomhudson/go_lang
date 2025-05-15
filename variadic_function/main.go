package main

import "fmt"

func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Hello", "World!", 2, 3, 4, 4, 5, 5)
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println("The sum is:", result)

}
