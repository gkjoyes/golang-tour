// Sample program to show how to declare and iterate over arrays of different types.
package main

import "fmt"

func main() {

	// Different type of declarations.
	declare1() //Declares an empty array and initialized it with zero values.
	declare2()
	declare3()
}

// Declares an empty array and initialized it with zero values.
func declare1() {

	// Declare an array of five strings that is initialized to its zero value.
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"

	// Iterate over the array of strings.
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}

// Using the literal style for declaring an array and initialized it with some values.
func declare2() {
	// Declare an array of 3 integers that is initialized with some values.
	numbers := [3]int{10, 20, 30}

	// Iterate over the array of numbers.
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
}

// Seeks the help of 3 dot operator for declaring and initializing an array.
func declare3() {
	// Declare an array of 3 floating point numbers that is initialized with some values.
	values := [...]float32{1.0, 2.0, 3.0}

	// Iterate over the array of floating point numbres.
	for i := 0; i < len(values); i++ {
		fmt.Println(i, values[i])
	}
}
