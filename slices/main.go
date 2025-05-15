package main

import "fmt"

// slice -> dynamic array
// most used construct in go
func main() {
	// uninitialized slice is nil
	var num []int
	fmt.Println(num)
	fmt.Println(num == nil)
	fmt.Println(len(num))
	// create a slice with make
	num = make([]int, 0, 10)
	fmt.Println(num)
	fmt.Println(num == nil)
	// capacity -> maximum numbers of elements can fit
	fmt.Println(cap(num))
	num = append(num, 1)
	num = append(num, 2)
	num = append(num, 3)
	num = append(num, 4)
	num = append(num, 5)
	num = append(num, 6)
	num = append(num, 7)
	num = append(num, 8)
	num = append(num, 9)
	num = append(num, 10)
	num = append(num, 11)
	num = append(num, 12)
	num = append(num, 13)
	fmt.Println(num)
	fmt.Println(cap(num))
	// copy slice
	num2 := make([]int, len(num))
	copy(num2, num)
	fmt.Println("nums2: ", num2)
	// slice of slice
	num3 := num[0:5]
	fmt.Println(num3)

}
