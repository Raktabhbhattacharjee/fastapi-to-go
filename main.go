// package main tells Go to build an executable, like a C++ program with main().
package main

import "fmt"

// Go imports packages directly; fmt provides printing functions, not the program's entry point.
// fmt.Println is like C++ std::cout or Python's print() for terminal output.
// main is the entry point, like C++ main(); Python often uses an if __name__ == "__main__": guard.
// Unlike Go's required main function, Python's __main__ guard is optional.
func main() {
	fmt.Println("Learn golang in such a way so you know how to build softwares in the age of ai era ")
}

// Conclusion: package main makes this an executable, main() starts it, and fmt.Println prints text.
