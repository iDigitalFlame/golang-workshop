package main

import "fmt"

/*
	Lab 2: Pointers and References

	Goal: Use pointers to modify values.
*/

func main() {

	Number1 := 1

	// Make PointerToNumber1 a pointer to the variable Number1.
	var PointerToNumber1 *int

	// Modify this to print the memory address of PointerToNumber1.
	fmt.Printf("", PointerToNumber1)

	// Change
	AnotherPointerToNumber1 := PointerToNumber1
	// Does this change the value of Number1?
	*AnotherPointerToNumber1 = 5

	// Edit this to print the value of that AnotherPointerToNumber1 points to
	fmt.Printf("", AnotherPointerToNumber1)

	// Modify this to print the value of Number1.
	fmt.Printf("", Number1)
}
