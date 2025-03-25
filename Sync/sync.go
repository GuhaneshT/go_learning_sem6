package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	counter int
)

func incrementer(mu *sync.Mutex, counter *int) {
	mu.Lock()
	*counter += 1  // Corrected increment syntax
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() { //anonymous goroutine captures variables from the outer scope and executes on creation
			defer wg.Done() // Mark goroutine as done when it exits
			incrementer(&mu, &counter)
		}()
	}
	wg.Wait()
	 // Wait for all goroutines to finish before printing
	fmt.Println("Counter:", counter) // Corrected print statement
	
}
