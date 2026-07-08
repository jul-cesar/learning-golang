package main

import (
	"fmt"
)

func main() {

	// Arrays 
	// the firt parameter is the length of the array and the second parameter is the type of the array
	numbers := [5]int{1, 2, 3, 4, 5}
	// to print an array we can use the %v verb in fmt.Printf
	fmt.Printf("This is the array %v\n", numbers)
	// with len() function we can get the length of an array
	fmt.Println("The array length is", len(numbers))

	// multidimensional arrays
	// the first parameter is the length of the array and the second parameter is the length of the inner array and the third parameter is the type of the array+
	multi := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("This is the multidimensional array %v\n", multi)
}
