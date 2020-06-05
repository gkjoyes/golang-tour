// Sample program to show how maps behave when you read an absent key.
package main

import "fmt"

func main() {

	// Create a map to track scores for players in the game.
	scores := make(map[string]int)

	// Read the element at key "tony". It is absent so we get the zero-value for this map's value type.
	score := scores["tony"]

	fmt.Println("Score:", score)

	// If we need to check for the presence of a key we use a 2 variable assignment. The 2nd variable is a bool.
	score, ok := scores["tony"]

	fmt.Println("Score: ", score, "Present: ", ok)

	// We can leverage the zero-value behaviour to write convenient code like this:
	scores["tony"]++

	// Without this behaviour, we would have to code in a defensive way like this.
	if n, ok := scores["tony"]; ok {
		scores["tony"] = n + 1
	} else {
		scores["tony"] = 1
	}

	score, ok = scores["tony"]
	fmt.Println("Score:", score, "Present:", ok)
}
