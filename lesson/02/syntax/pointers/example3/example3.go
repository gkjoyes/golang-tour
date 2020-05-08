// Sample program to show the basic concept of using a pointer to share data.
package main

import "fmt"

// user represents a user in the system.
type user struct {
	name   string
	email  string
	logins int
}

func main() {

	// Declare and initialize a variable named "u1" of type user.
	u1 := user{
		name:  "user",
		email: "user@email.com",
	}

	// Pass the address of "u1" value.
	display(&u1)

	// Pass the address of the logins field from within the u1 value.
	increment(&u1.logins)

	// Pass the address of "u1" value.
	display(&u1)
}

// display declares "u" as a pointer variable of type user whose value is always an address.
func display(u *user) {
	fmt.Printf("%p\t%+v\n", u, *u)
}

// increment declares logins as a pointer variable of type int whose value is always an address.
func increment(logins *int) {
	*logins++
}
