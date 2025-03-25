package main

import (
	"fmt"
)

// Task structure to store computation requests
type Task struct {
	i, j int
}

// Function to multiply matrices using a worker pool and buffered channel
func multiplyMatrices(A, B [][]int, workers int) [][]int {
	rowsA, colsA := len(A), len(A[0])
	colsB := len(B[0])

	// Initialize the result matrix
	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	// Buffered channel to manage worker tasks
	taskChan := make(chan Task, workers)
	doneChan := make(chan struct{})

	// Worker function
	worker := func() {
		for task := range taskChan {
			sum := 0
			for k := 0; k < colsA; k++ {
				sum += A[task.i][k] * B[k][task.j]
			}
			result[task.i][task.j] = sum
		}
		doneChan <- struct{}{} // Signal completion
	}

	// Start worker pool
	for w := 0; w < workers; w++ {
		go worker()
	}

	// Distribute tasks
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			taskChan <- Task{i, j}
		}
	}
	close(taskChan) // No more tasks to add

	// Wait for workers to finish
	for w := 0; w < workers; w++ {
		<-doneChan
	}

	return result
}

func main() {
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	B := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	fmt.Println("Matrix A:")
	for _, row := range A {
		fmt.Println(row)
	}

	fmt.Println("\nMatrix B:")
	for _, row := range B {
		fmt.Println(row)
	}

	workers := 4 // Number of workers (can be adjusted)
	result := multiplyMatrices(A, B, workers)

	fmt.Println("\nResultant Matrix:")
	for _, row := range result {
		fmt.Println(row)
	}
}
