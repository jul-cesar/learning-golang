package main

import "fmt"

func main() {
	//variables
	// var keyword is used to declare a variable
	// name of the variable is name and the type of the variable is string
	var name string = "Julio Cesar"
	city := "Mexico City" // short variable declaration, type is inferred

	fmt.Printf("Hello, my name is %s and I live in %s.\n", name, city)
	// A way to declare multiple variables at once is to use a var block
	var (
		age       int     = 22
		height    float64 = 1.75
		isStudent bool    = true
	)

	fmt.Printf("im %d, mi height is %f and is %t im student", age, height, isStudent)

	// Arguemnt formats
	// %d - integer
	// %f - float
	// %s - string
	// %t - boolean

	// Default values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("Default values: int=%d, float=%f, string='%s', bool=%t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	//constants

	const pi = 3.14
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
	)

	// Constant can be untyped or typed. Untyped constants are more flexible and can be used in different contexts, while typed constants have a specific type and cannot be used in contexts that require a different type.

	// Also, the compiler doesnt bother when a constant is declared but not used, but it will throw an error if a variable is declared but not used.
}
