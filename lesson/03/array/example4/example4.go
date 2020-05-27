// Sample program to show how the for-range has both value and pointer semantics.
package main

import "fmt"

func main() {

	// Using the pointer semantic form of the for-range.
	friends := [5]string{"A", "B", "C", "D", "E"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i := range friends {
		friends[1] = "J"
		if i == 1 {
			fmt.Printf("Aft[%s]\n", friends[1])
		}
	}

	// Using the value semantic form of the for-range.
	friends = [5]string{"A", "B", "C", "D", "E"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range friends {
		friends[1] = "J"
		if i == 1 {
			fmt.Printf("Aft[%s]\n", v)
		}
	}

	// Using the value semantic form of the for-range but with pointer semantic access.
	friends = [5]string{"A", "B", "C", "D", "E"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range &friends {
		friends[1] = "J"

		if i == 1 {
			fmt.Printf("Aft[%s]\n", v)
		}
	}
}
