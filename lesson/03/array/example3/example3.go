// Sample program to show how the behaviour of the for range and how memory for an array is contiguous.
package main

import "fmt"

func main() {

	// Declare an array of 5 strings initialized with some values.
	friends := [5]string{"A", "B", "C", "D", "E"}

	// Iterate over the array and display the value and address of each element.
	for i, v := range friends {
		fmt.Printf("Value[%s]\tAddress[%p]\tIndexAddr[%p]\n", v, &v, &friends[i])
	}
}
