package main

import (
	"fmt"
	"sync"
)

func pingPong(sendCh, recvCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		// Receive message
		msg := <-recvCh
		fmt.Println(msg)

		// Send next message
		if msg == "Ping" {
			sendCh <- "Pong"
		} else {
			sendCh <- "Ping"
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ping := make(chan string)
	pong := make(chan string)

	wg.Add(1)

	// Start Goroutine that receives from ping channel and sends to pong channel
	go pingPong(pong, ping, &wg)

	// Start the game by sending "Ping"
	ping <- "Ping"

	// Main goroutine plays the other side of the game
	for i := 1; i <= 10; i++ {
		// Receive response
		response := <-pong
		fmt.Println(response)

		// Send next message if not at the end
		if i < 10 {
			ping <- "Ping"
		}
	}

	// Wait for Goroutine to finish
	wg.Wait()

	// Close the channels
	close(ping)
	close(pong)

	fmt.Println("Game Over!")
}