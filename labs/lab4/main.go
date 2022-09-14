package main

import "fmt"

/*
	Lab 4: Maps

	Goal: Understand maps how to create them.
*/

func main() {

	Map1 := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
	}

	// Modify this to print the value stored with key "3".
	fmt.Printf("%s\n", Map1["3"])

	// Add this value to Map1 with the key "4".
	v := "four"
	Map1["4"] = v

	// Modify this to print the length of the Map
	fmt.Printf("%d\n", len(Map1))

	var Map2 map[string]string = nil
	Map2 = make(map[string]string)

	// Is this valid?
	Map2["two"] = "2"

	// Initialize a new empty Map
	var Map3 map[string]string = make(map[string]string)

	/// Is this valid?
	delete(Map3, "This key does not exist")

	// Is this valid? If so, print the value of Result1 as a string.
	Result1 := Map3["non-exist"]

	// Is this valid? If so, what is the value of Result2 and ok, and then print them!
	Result2, ok := Map3["I-dont-exist"]

	fmt.Printf("result1 : %s, %t\n", Result2, ok)

}
