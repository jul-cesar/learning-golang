package main

import (
	"fmt"
)


func main() {
	// For loop that iterates 5 times
	for i := 0; i < 5; i++ {
		fmt.Println("This is iteration number:", i)
	}

	// In goglang theres no while loop, but you can use a for loop to achieve the same functionality
	counter := 0
	for counter < 5 {
		fmt.Println("Counter is at:", counter)
		counter++
	}
	// Infinite loop
	for {
		break // This will break the infinite loop immediately
	}

	

}
