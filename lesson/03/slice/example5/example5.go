// Sample program to show how the for-range has both value and pointer semantics.
package main

import "fmt"

func main() {

	// Using the value semantic form of the for-range.
	friends := []string{"A", "B", "C", "D", "E"}
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", v)
	}

	// Using the pointer semantic form of the for-range.
	friends = []string{"A", "B", "C", "D", "E"}
	for i := range friends {
		fmt.Println(len(friends))
		friends = friends[:2]
		fmt.Printf("v[%s]\n", friends[i])
	}
}
