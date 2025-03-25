package basics

import (
	"fmt"
)

//regular function
func add_numbers(a int,b int) int {
	return a+b
}
//named return values
func add_numbers_named(a int,b int) (sum int) {
	sum = a+b
	return
}
//multiple return values
func add_numbers_multiple(a int,b int) (int,int) {
	return a+b,a-b
}
//variadic functions
func add_numbers_variadic(a ...int) int {
	sum := 0
	for _,val := range a {
		sum += val
	}
	return sum
}
//function as a variable
var add = func(a int,b int) int {
	return a+b
}
//function as a parameter
func apply_function(f func(int,int) int,a int,b int) int {
	return f(a,b)
}
//closure
func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
//defer
func defer_example() {
	defer fmt.Println("world")
	//deferred functions are executed in LIFO order
	fmt.Println("hello")
}
//panic and recover
func panic_example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()
	panic("Panic")
}
//structs
type person struct {
	name string
	age int
}
//methods
func (p person) get_name() string {
	return p.name
}
//interfaces
type animal interface {
	speak() string
}
type dog struct {
	name string
}
func (d dog) speak() string {
	return "woof"
}
//goroutines
func goroutine() {
	fmt.Println("goroutine")
}
//channels
func channel() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}
//select
func select_example() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 2
	}()
	select {
	case val := <-ch1:
		fmt.Println(val)
	case val := <-ch2:
		fmt.Println(val)
	}
}


func main(){
	// basic print statement
	fmt.Println("Hello World")
	//if else
	var i int = 10 // no need to use var again, just use i := 10 for short declaration - dynamic typing
	if i >= 10 {
		fmt.Println("i is greater than 10")
	} else {
		fmt.Println("i is less than 10")
	} // else should be on the same line as the closing brace of if

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	var arr [3]int = [3]int{1, 2, 3} //array decalaration

	for index,element := range arr { // range returns index and element
		fmt.Println(index, element)
	}

	// switch case
	i=2 //:= can be used only for new variables
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	default:
		fmt.Println("i is neither 1 nor 2")
	}

	//functions
	fmt.Println(add_numbers(1,2))
	fmt.Println(add_numbers_named(1,2))
	fmt.Println(add_numbers_multiple(1,2))
	fmt.Println(add_numbers_variadic(1,2,3,4,5))
	fmt.Println(add(1,2))
	fmt.Println(apply_function(add,1,2))
	c := closure()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	defer_example()
	panic_example()
	Person := person{name: "John", age: 25}
	fmt.Println(Person.get_name())
	d := dog{name: "Tom"}
	fmt.Println(d.speak())
	go goroutine()
	channel()
	select_example()
	
	


}