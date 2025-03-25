package main

import (
	"fmt"
	"sort"
)

type Student struct{
	Names [4]string
	Marks [4]int
}

func main(){
	arr := [5]int{1, 3, 2, 4, 5}

	student:=Student{
		Names: [4]string{"a","b","c","d"},
		Marks: [4]int{1,3,2,4},
	}
	fmt.Println(student)
	slice:=arr[:]

	sort.Ints(slice)

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	sort.Slice(student.Marks[:], func(i, j int) bool {
		return student.Marks[i] < student.Marks[j]
	})
	fmt.Println(student)

	
	fmt.Println("array sorted throug creating slice of array")
	fmt.Println(slice)
	fmt.Println("array sorted through sort.Ints")
	fmt.Println(arr)

	

}