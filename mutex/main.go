package main

import (
	"fmt"
	"sync"
)

/* What is Race condition?
When multiple resources or multiple process are trying to modify the same resource at the same time.
This can lead to unexpected behavior and bugs in the code.
or which leads to non atomicity of the resource.
*/

type post struct {
	views int
	mu    sync.Mutex // Mutex to protect the views field
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()     // Decrement the WaitGroup counter when the goroutine completes
		p.mu.Unlock() // Unlock the mutex after accessing the views field
	}()

	// Lock the mutex before accessing the views field
	// This ensures that only one goroutine can access the views field at a time
	p.mu.Lock() // Lock the mutex before accessing the views field
	p.views++

}

func main() {
	var wg sync.WaitGroup

	myPost := post{views: 0}
	for i := 0; i < 1000; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go myPost.increment(&wg)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Views:", myPost.views)

	/* 	About above code: Before Adding Mutex
	Race Condition: Running the above code will give different output every time.
	1. 	[Running] go run "go_lang\mutex\main.go"
			Views: 994
		[Done] exited with code=0 in 1.861 seconds

	2.	[Running] go run "go_lang\mutex\main.go"
			Views: 991
		[Done] exited with code=0 in 2.033 seconds
	*/
	// -------------------------------------------------------------
	/* 	About above code: After Adding Mutex
	1. 	[Running] go run "go_lang\mutex\main.go"
			Views: 1000
		[Done] exited with code=0 in 1.805 seconds

	2. 	[Running] go run "go_lang\mutex\main.go"
			Views: 1000
		[Done] exited with code=0 in 0.855 seconds

	3. 	[Running] go run "go_lang\mutex\main.go"
			Views: 1000
		[Done] exited with code=0 in 0.887 seconds
	*/
}
