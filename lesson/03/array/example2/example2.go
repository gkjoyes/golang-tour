// Sample program to show how arrays of different sizes are not of the same type.
package main

import "fmt"

func main() {

	// Declare an array of 4 integers that are initialized with some values.
	var five [5]int

	// Declare an array of 4 integers that are initialized with some values.
	four := [4]int{10, 20, 30, 40}

	// Assign one array to the other.
	// five = four (Error: cannot use four (type [4]int) as type [5]int in assignment)

	fmt.Println(five)
	fmt.Println(four)
}
