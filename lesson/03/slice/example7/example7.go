// Sample program to show how to declare and use variadic function.
package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	id   int
	name string
}

func main() {

	// Declare and initialize values of type user.
	u1 := user{
		id:   1,
		name: "George",
	}

	u2 := user{
		id:   2,
		name: "Cyril",
	}

	// Display both values.
	fmt.Println(u1, u2)

	// Create a slice of user values.
	u3 := []user{
		{id: 3, name: "Tony"},
		{id: 4, name: "Jishnu"},
	}

	// Display all the user values from the slice.
	display(u3...)

	// Do some modifications on slice.
	change(u3...)

	// Display again after the modification.
	display(u3...)

}

// display can accept and display multiple values of user types.
func display(users ...user) {
	fmt.Printf("\n")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

// change shows how the backing array is shared.
func change(users ...user) {
	users[1] = user{id: 0, name: "CHANGE"}
}
