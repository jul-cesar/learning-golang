package main

import (
	"fmt"
)

func main() {

	// Maps

	// Maps in golang are unordered collections of key-value pairs. They are similar to dictionaries in Python or hash tables in other programming languages. Maps are useful for storing and retrieving data based on unique keys.

	// to create a map we can use the map keyword
	// The argument to the map keyword is the type of the key and the type of the value. For example, to create a map that maps strings to integers, we can use the following syntax:
	// myMap := map[string]int{})
	CapitalCities := map[string]string{
		"France":  "Paris",
		"Germany": "Berlin",
		"Italy":   "Rome",
	}
	fmt.Printf("Capital cities: %v\n", CapitalCities)

	// To retrieve a value from a map, we can use the key as an index. For example, to get the capital city of France, we can use the following syntax:
	capital := CapitalCities["France"]
	fmt.Printf("The capital of France is: %s\n", capital)

	// If the value doesnt not exist, the zero value of the value type will be returned. For example, if we try to get the capital city of Spain, which is not in the map, we will get an empty string:
	capital = CapitalCities["Spain"]
	fmt.Printf("The capital of Spain is: %s\n", capital)

	// To check if a key exists in a map, we can use the two-value assignment syntax. The second value will be a boolean indicating whether the key exists in the map. For example:
	capital, ok := CapitalCities["Spain"]

	if ok {
		fmt.Printf("The capital of Spain is: %s\n", capital)
	} else {
		fmt.Println("Spain is not in the map")
	}

	// to delete a key-value pair from a map, we can use the delete function. For example, to delete the capital city of Germany from the map, we can use the following syntax:

	delete(CapitalCities, "Germany")
}
