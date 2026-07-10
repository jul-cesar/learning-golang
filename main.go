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

	fmt.Println("Name before modification:", person.Name)
	// Pass the address of the person struct to the modifyPerson function to modify its Name field. Using & allows us to pass a pointer to the struct, so we can modify the original struct instead of a copy.
	person.modifyPerson()
	fmt.Println("Name after modification:", person.Name)

	x := 10
	ptr := &x // Create a pointer to the variable x
	fmt.Println("Value of x>: %t and address: %p", x, ptr)

	*ptr = 20 // Modify the value of x using the pointer
	fmt.Println("Value of x after modification using pointer:", x) 
}

// The modifyPerson function takes a pointer to a Person struct as an argument. It modifies the Name field of the struct and prints the modified value.
// func modifyPerson(person *Person) {
// 	person.Name = "Modified Name"
// 	fmt.Println("Name after modification inside function:", person.Name)
// }

// Instead of defining the modifyPerson function as a standalone function, we can define it as a method on the Person struct. This allows us to call the method directly on an instance of the Person struct, making the code more readable and organized.
func (person *Person) modifyPerson() {
	person.Name = "Modified Name"
	fmt.Println("Name after modification inside function:", person.Name)
}
