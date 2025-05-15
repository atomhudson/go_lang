package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		// fmt.Println(nums[i])
	}
	sum := 0
	// range
	for _, v := range nums {
		sum += v
		fmt.Println(sum)
	}
	fmt.Println(sum)
	// range with index
	for i, j := range "hello" {
		println(i, j)
	}

	for k, v := range map[string]int{"a": 1, "b": 2} {
		println(k, v)
	}
}
