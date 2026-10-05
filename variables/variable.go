
package main

import "fmt"

// PACKAGE-LEVEL SCOPE (Outer Scope)
// Constants and variables declared here are accessible anywhere in the package.
const a = "hello"

func main() {
	// CONCEPT: VARIABLE SHADOWING
	// Declaring 'a' inside main() "shadows" (hides) the package-level 'a' above.
	// Any reference to 'a' inside main() will now resolve to "initial", NOT "hello".
	// Unlike Python (which has globals()), Go has no built-in syntax to access a 
	// shadowed variable in the same scope. Standard idiom: give them unique names!
	var a = "initial"
	fmt.Println(a) // Prints: initial

	// MULTIPLE VARIABLE DECLARATION
	// Go allows declaring and initializing multiple variables of the same type in one line.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// TYPE INFERENCE
	// Go infers the type (bool) automatically based on the assigned value.
	var d = true
	fmt.Println(d)

	// ZERO VALUES
	// Variables declared without an explicit initial value automatically get 
	// their type's "zero value" (0 for int, "" for string, false for bool, nil for pointers).
	var e int
	fmt.Println(e) // Prints: 0

	// SHORT VARIABLE DECLARATION operator (:=)
	// Shorthand for declaring and initializing variables inside functions.
	// Go automatically infers string type. Note: := cannot be used at package level.
	f := "apple"
	fmt.Println(f)
	
	// Refers to the shadowed local variable 'a'
	fmt.Println(a) 

	constExample()
}
// shorthand syntax 
// name:="raktabh"
// how to declare a variable 
