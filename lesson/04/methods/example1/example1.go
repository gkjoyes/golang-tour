// Sample program to show how to declare methods.
package main

import (
	"fmt"
	"strings"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending email to %s <%s>\n", u.name, u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	// Value of type user can be used to call methods declared with both value and pointer receivers.
	george := user{name: "george", email: "george@gmail.com"}
	george.changeEmail("george@hotmail.com")
	george.notify()

	// Pointers of type user can also be used to call methods declared with both value and pointer receiver.
	tony := user{name: "tony", email: "tony@gmail.com"}
	tony.changeEmail("tony@hotmail.com")
	tony.notify()

	// Create a slice of user values with two users.
	users := []user{
		{name: "manu", email: "manu@gmail.com"},
		{name: "cyril", email: "cyril@gmail.com"},
	}

	// Iterate over the slice of users switching semantics. Not Good!
	for _, u := range users {
		// This is not going to work, because we are copying each user to variable u.
		u.changeEmail(strings.Replace(u.email, "gmail", "hotmail", 1))
	}

	fmt.Println(users)
}
