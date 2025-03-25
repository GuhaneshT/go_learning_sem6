package main

import (
	"fmt"
	"sort"
)

// Function to demonstrate arrays
func arrayExample() {
	// Declare and initialize an array
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}

	// Print the array and its length
	fmt.Println("Array:", numbers)
	fmt.Println("Length of array:", len(numbers))

	// Iterate over the array
	for i, v := range numbers {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}

// Function to demonstrate slices
func sliceExample() {
	// Create a slice
	numbers := []int{5, 3, 4, 1, 2}

	// Append to the slice
	numbers = append(numbers, 6, 7)

	// Print the slice, its length, and capacity
	fmt.Println("Slice:", numbers)
	fmt.Println("Length of slice:", len(numbers))
	fmt.Println("Capacity of slice:", cap(numbers))

	// Sort the slice
	sort.Ints(numbers)
	fmt.Println("Sorted Slice:", numbers)
}

// Function to demonstrate maps
func mapExample() {
	// Create a map
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25

	// Print the map and its length
	fmt.Println("Ages:", ages)
	fmt.Println("Length of map:", len(ages))

	// Check existence of a key
	if age, exists := ages["Alice"]; exists {
		fmt.Println("Alice's age:", age)
	}

	// Delete a key
	delete(ages, "Bob")
	fmt.Println("Ages after deletion:", ages)

	// Iterate over the map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}

// Function to demonstrate custom sorting
func customSortExample() {
	// Create a slice of strings
	strings := []string{"apple", "banana", "kiwi", "grape"}

	// Sort strings by length
	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) < len(strings[j])
	})

	// Print sorted strings
	fmt.Println("Sorted by length:", strings)
}

func main() {
	fmt.Println("Array Example:")
	arrayExample()
	fmt.Println()

	fmt.Println("Slice Example:")
	sliceExample()
	fmt.Println()

	fmt.Println("Map Example:")
	mapExample()
	fmt.Println()

	fmt.Println("Custom Sort Example:")
	customSortExample()
}