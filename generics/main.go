package main

import "fmt"

// func printSlice[T int | string | bool | float32](item []T) { // both function siganture works
func printSlice[T comparable, V string](item []T, v V) { // both function siganture works
	for _, i := range item {
		fmt.Print(i, " ", v, " ")
	}
	fmt.Println()
}

type stack[T any] struct {
	elements []T
}

func main() {
	printSlice([]int{1, 2, 3, 4, 5}, "int")
	printSlice([]string{"a", "b", "c", "d", "e"}, "string")
	printSlice([]float32{1.1, 2.2, 3.3, 4.4, 5.5}, "float32")
	printSlice([]bool{true, false, true, false, true}, "bool")

	stack := stack[int]{
		elements: []int{1, 2, 3},
	}

	fmt.Println("Stack elements:", stack.elements)
}
