// Sample program to teach the mechanics of escape analysis.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

func main() {
	u1 := createUser1()
	u2 := createUser2()

	println("u1", &u1, "u2", u2)
}

// createUser1 creates a user value and pass a copy back to the caller.
//go:noinline
func createUser1() user {
	u := user{
		name:  "user1",
		email: "user1@email.com",
	}

	println("u1", &u)
	return u
}

// createUser2 creates a user value and shares the value with the caller.
//go:noinline
func createUser2() *user {
	u := user{
		name:  "user2",
		email: "user2@email.com",
	}

	println("u2", &u)
	return &u
}

// To See escape analysis and inlining decisions.
// go build -gcflags -m=2
