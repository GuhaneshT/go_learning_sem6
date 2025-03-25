package main

import (
	"fmt"
	"sync"
)

// Function to generate Fibonacci sequence
func fibonacci(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark goroutine as done when it exits

	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a // Send Fibonacci number
		a, b = b, a+b
	}
	close(ch) // Close channel after sending all numbers
}

func main() {
	n := 10 // Number of Fibonacci numbers to generate
	ch := make(chan int) // Create a channel for communication
	var wg sync.WaitGroup

	wg.Add(1)
	go fibonacci(n, ch, &wg) // Start Fibonacci goroutine

	// Read from channel and print values
	for num := range ch {
		fmt.Println(num)
	}

	wg.Wait() // Ensure all goroutines finish
	fmt.Println("Done!")
}
