package main

import "fmt"

func main() {
	// Declare an integer variable
	x := 10

	// Declare a pointer to an integer
	var p *int

	// Assign the address of x to p
	p = &x

	// Print the initial values
	fmt.Printf("Initial value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of p: %p\n", p)
	fmt.Printf("Value pointed to by p: %d\n", *p)

	// Modify the value through the pointer
	*p = 20

	// Print the updated values
	fmt.Printf("\nAfter modifying through pointer:\n")
	fmt.Printf("New value of x: %d\n", x)
	fmt.Printf("Value pointed to by p: %d\n", *p)

	// Modify the value directly
	x = 30

	// Print the final values
	fmt.Printf("\nAfter modifying x directly:\n")
	fmt.Printf("Final value of x: %d\n", x)
	fmt.Printf("Value pointed to by p: %d\n", *p)
}
