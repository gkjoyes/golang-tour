// Sample program to show how maps are reference types.
package main

import "fmt"

func main() {

	// Initialize a map with values.
	scores := map[string]int{
		"george": 25,
		"tony":   20,
	}

	// Pass the map to a function to perfrom some mutation.
	double(scores, "george")

	// See the change is visible in our map.
	fmt.Println("Score:", scores["george"])

}

// double finds the score for a specific palyer and multiplies it by 2.
func double(scores map[string]int, player string) {
	scores[player] = 2 * scores[player]
}
