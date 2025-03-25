package main

import (
	"fmt"
	"sync"
)

var(
	mu sync.Mutex
	wg sync.WaitGroup
	sum int
)

func add(num ...int) {
defer wg.Done()
for _,value := range num{
	mu.Lock()
	sum+=value 
	mu.Unlock()
}
}

func main(){
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	first_half_of_x := arr[:len(arr)/2]
	second_half_of_x := arr[len(arr)/2:]
	fmt.Println("First half of array:", first_half_of_x)
	fmt.Println("Second half of array:", second_half_of_x)
	wg.Add(2)
	go add(first_half_of_x...)
	go add(second_half_of_x...)
	wg.Wait()
	fmt.Println("the sum is ")
	fmt.Println(sum)

}