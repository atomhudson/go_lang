package main

import "fmt"

func main() {
	// int
	fmt.Println(1 + 1)
	// String
	fmt.Println("Hello " + "World")
	// bool
	fmt.Println(true && false)
	// float
	fmt.Println(1.5 + 2.5)
	// complex
	fmt.Println(1 + 2i)
	// rune
	fmt.Println('a')
	// byte
	fmt.Println(byte('a'))
	// array
	arr := [3]int{1, 2, 3}
	arr[0] = 3
	fmt.Println(arr)
	// slice
	slice := []int{1, 2, 3}
	slice2 := make([]int, 3)
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	fmt.Println(slice2)
	fmt.Println(slice)
	// map
	m := map[string]int{"a": 1, "b": 2}
	fmt.Println(m)
}
