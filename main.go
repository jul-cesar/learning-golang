/* every Go program starts with a package declaration. The package name is the first line of every Go source file. The package main is obligatory for executable programs, the compler will search for the main function */
package main

// this is the syntax of an import statement, it is used to import packages from the standard library or from other packages. Something important abount import in the main package is that it cant import anything from the project but it cant be imported by other packages //
import (
	"fmt"
)

/* this is the syntax of a function, with the func keyword, very similar to other languages like TypeScript */

// the main function is the entry point of the program, it is called when the program starts executing, in this case we use it to print Hello world using de fmt (format) package from Go //
func main() {
	fmt.Println("Hello world")
}
