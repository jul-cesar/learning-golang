package main

import (
	"fmt"
)

func main() {

	// slices

	numbers := []int{1, 2, 3, 4, 5}
	// this is how you can create a slice from an array or another slice
	allNumbers := numbers[:]
	fmt.Printf("All numbers: %v\n", allNumbers)
	// this takes the first three numbers from the slice
	firstThreeNumbers := numbers[0:3]
	fmt.Printf("First three numbers: %v\n", firstThreeNumbers)

	// this is how you can create new slices
	fruits := []string{
		"apple",
		"banana",
		"cherry",
	}
	fmt.Printf("Fruits: %v\n", fruits)

	// we cand append to a slice using the built-in append function
	fruits = append(fruits, "date")
	fmt.Printf("Fruits after appending: %v\n", fruits)

	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

}
