package main

import "fmt"

// numbered sequesnce of specific length
func main() {
	var nums [5]int
	// array length
	fmt.Println(len(nums))
	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(nums)
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	nums[4] = 5
	// print all elements
	for i := 0; i < len(nums); i++ {
		fmt.Print(fmt.Sprintf("%d ", nums[i]))
	}
	str := [3]string{"heloo", "world", "!"}
	// print all elements
	fmt.Println(str)
}
