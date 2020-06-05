// Sample program to shows how iterating over a map is random.
package main

import "fmt"

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

func main() {

	// Declare and initialize the map with values.
	users := map[string]user{
		"George": {name: "George", surname: "Joyes"},
		"Tony":   {name: "Tony", surname: "Tomy"},
		"Jishnu": {name: "Jishnu", surname: "TG"},
		"Cyril":  {name: "Cyril", surname: "George"},
	}

	// Iterate over the map and print each key and value.
	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Printf("\n")

	// Iterate over the map and just print the keys. Notice the results are different.
	for key := range users {
		fmt.Println(key)
	}
}
