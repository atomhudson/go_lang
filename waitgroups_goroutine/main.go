package main

import (
	"fmt"
	"sync"
)

/*
1. The **WaitGroup** is a synchronization primitive that allows you to wait for a collection of goroutines to finish.
2. It is used to wait for a collection of goroutines to finish executing.
--------------------------------------------------------Or--------------------------------------------------------
It ensures that all goroutines have completed before the main function exits.
*/
func task(id int, w *sync.WaitGroup) {
	/*
		**defer** statment is used to ensure that the **WaitGroup** counter is decremented when the **goroutine** completes
		or runs at the end of the function
	*/
	defer w.Done() // Decrement the counter when the goroutine completes
	fmt.Println("Task", id, "is running")
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go task(i, &wg)
	}
	wg.Wait() // Wait for all goroutines to finish
	// time.Sleep() // is not needed here because we are using WaitGroup to wait for all goroutines to finish
	// time.Sleep(2 * time.Second) // Wait for goroutines to finish
}
