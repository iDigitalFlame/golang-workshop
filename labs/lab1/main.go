package main

import "fmt"

/*
	Lab 1: Math Stuff

	Goal: Do some math on the following examples and print out their responses.
*/

func main() {

	// Add the values of Number1 and Number2 and store them in Number3.
	var Number1, Number2 int = 5, 10
	var Number3 int = Number1 + Number2

	// Edit this to Printf statement out the value of Number3 correctly.
	fmt.Printf("%d\n", Number3)

	// Subtract the values of Number4 and Number5 and store them in Number6.
	var Number4, Number5 uint16 = 15, 5
	var Number6 uint8
	Number6 = uint8(Number4 - Number5)

	// Edit this to Printf statement out the value of Number6 as a lowercase hex value.
	fmt.Printf("%X\n", Number6)

	// Multiple the values of Number4 and Number5 and store them in Number6.
	var Number7 float32 = 2.5
	var Number8 int = 50
	var Number9 float64 = float64(Number8) * float64(Number7)

	// Edit this to Printf statement out the value of Number9 correctly.
	fmt.Printf("%.2f\n", Number9)
}
