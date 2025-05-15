package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func addAndReturnMultiple() (string, string, string, int) {
	return "golang", "javascript", "java", 8
}

func processIT(fn func(a int) int) {
	fn(5)
}

func main() {
	// This is a placeholder for the main function.
	// You can add your code here to implement the desired functionality.
	// For example, you might want to call other functions or perform some operations.
	// Example: Call a function to perform a specific task

	result := add(5, 3)
	println("The result of adding 5 and 3 is:", result)
	lang1, lang2, lang3, num := addAndReturnMultiple()
	fmt.Println("Language 1:", lang1)
	fmt.Println("Language 2:", lang2)
	fmt.Println("Language 3:", lang3)
	fmt.Println("Number:", num)

	fn := func(a int) int {
		return a * 2
	}
	processIT(fn)

}
