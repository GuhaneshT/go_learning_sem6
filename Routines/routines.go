package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	wg sync.WaitGroup
)
func one_to_ten_a(){
	for i := 1; i<=10; i++ {
		fmt.Println("from a: " + fmt.Sprintf("%d", i))
	}
	wg.Done()

}

func one_to_ten_b(){
	for i := 1; i<=10; i++ {
		fmt.Println("from b: " + fmt.Sprintf("%d", i))
	}
	wg.Done()

}
func main(){
	wg.Add(2)
	go one_to_ten_a()
	go one_to_ten_b()
	wg.Wait()
	fmt.Println("Done")

}
