package main

import "fmt"

func main() {
	var myInt int

	yyy := 0

	yyy = yyy * 5
	yyy *= 5
	_ = yyy

	myInt = 5
	myInt *= 100

	fmt.Printf("Here is the value of myInt: %x\n", myInt)

	myString := "This is a string"

	fmt.Println(myString)

	var g bool
	var t float32

	var myString2 string
	_ = myString2

}
