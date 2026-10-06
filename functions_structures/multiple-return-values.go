package main

import "fmt"

// This function returns two values:
// the name of a person and their age.
func getPerson() (string, int) {
	return "Rishi", 21
}

func main() {

	// getPerson() returns two values.
	// We store both values in name and age.
	name, age := getPerson()

	fmt.Println(name)
	fmt.Println(age)

	// Here we only care about the age.
	// The first returned value (name) is ignored using _.
	_, personAge := getPerson()

	fmt.Println(personAge)
}