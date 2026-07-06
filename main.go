package main

func main() {

	resultAdd := Add(2, 3)
	println(resultAdd)

	resultSum, resultProduct := SumAndProduct(4, 5)
	println(resultSum)
	println(resultProduct)

}

// To declare a function simply use the func keyword followed by the function name, parameters, and return type. The function body is enclosed in curly braces.
// Type can be omitted if it can be inferred from the context. For example, in the SumAndProduct function, the types of a and b are both int, so we can omit the type for b.

func Add(a int, b int) int {
	return a + b
}

// The SumAndProduct function takes two integers as input and returns their sum and product as two separate return values. The return types are specified in parentheses after the parameter list.
func SumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

// Also the capitalize name of the function indicates that it is exported and can be accessed from other packages. If the function name starts with a lowercase letter, it is unexported and can only be accessed within the same package.
