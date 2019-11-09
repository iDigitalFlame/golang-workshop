package main

import "fmt"

/*
	Lab 3: Arrays and Slices

	Goal: Understand arrays and slices and how to create them.
*/

func main() {

	Slice1 := []string{"one", "two", "three"}

	// Modify this to only print out the second element in the slice.
	fmt.Printf("", Slice1)

	// Modify this to create a slice with a size of 10.
	var Slice2 []string

	// Modify this to print the length of Slice2.
	fmt.Printf("", Slice2)

	var Slice3 []string = nil
	Slice3 = append(Slice3, "Is this valid?")

	Slice4 := make([]string, 0, 10)
	Slice4[5] = "Is this valid?"

	Slice5 := make([]string, 10, 20)
	Slice5[5] = "Is this valid?"

	var Array1 [4]string
	Array1[1] = "Is this valid?"

	// Modify this to only print out the second element in the array.
	fmt.Printf("", Array1)

	// Modify this to work. (hint: Array1 must be "converted").
	Result1 := append(Array1, "Is this valid?")

	// Modify this to print the length of Result1
	fmt.Printf("", Result1)

}
