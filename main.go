package main

import "fmt"

// The way you declare a struct in Go is by using the type keyword followed by the name of the struct and the struct fields enclosed in curly braces.

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// Nested structs are structs that contain other structs as fields. In this example, the Person struct contains an Address struct as one of its fields.
type Person struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	// Create an instance of the Person struct and assign values to its fields.
	person := Person{
		Name: "Julio",
		Age:  30,
	}

	fmt.Printf("This is a person struct: %+v\n", person)

	// Anonymous structs are useful when you need a struct for a specific purpose and don't want to define a new type.
	anonymousStruct := struct {
		height int
		weight int
	}{
		height: 180,
		weight: 75,
	}
	fmt.Printf("This is an anonymous struct: %+v\n", anonymousStruct)

	

}
