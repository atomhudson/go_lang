package main

import "fmt"

// This is a simple program to demonstrate the use of pointers in Go.
func changeNum(num int) {
	num = 10
	fmt.Println("Inside changeNum:", num)

}

// The changeNum function takes an integer as an argument and changes its value.
func changeNumPointer(num *int) {
	*num = 10
	fmt.Println("Inside changeNumPointer:", *num)
}

func main() {
	num := 1
	fmt.Println("Initial value of num:", num)
	changeNum(num)
	fmt.Println("Memmory address of num:", &num)
	fmt.Println("Value of num:", num)

	changeNumPointer(&num)
	fmt.Println("After changeNumPointer:", num)

}
