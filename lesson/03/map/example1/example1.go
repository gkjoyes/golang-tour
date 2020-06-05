// Sample program to show how to initialize a map, write to it, then read and delete from it.
package main

import "fmt"

// user represents someone using this program.
type user struct {
	name    string
	surname string
}

func main() {

	// Declare and make a map that stores values of type user with a key of type string.
	users := make(map[string]user)

	// 1. Add key/value pairs to the map.
	users["George"] = user{name: "George", surname: "Joyes"}
	users["Tony"] = user{name: "Tony", surname: "Tomy"}
	users["Cyril"] = user{name: "Cyril", surname: "George"}
	users["Jishnu"] = user{name: "Jishnu", surname: "TG"}

	// 2. Read the value at a specific key.
	tony := users["Tony"]
	fmt.Printf("%+v\n", tony)

	// 3. Replace the value at the "Tony" key.
	users["Tony"] = user{name: "Tomy", surname: "Tony"}

	// Read the "Tony" key again.
	fmt.Printf("%+v\n", users["Tony"])

	// 4. Delete the value at a specific key.
	delete(users, "Tony")

	// Check the length of the map. There are only 3 elements.
	fmt.Println(len(users))

	// It is safe to delete an absent key.
	delete(users, "Tony")
}
