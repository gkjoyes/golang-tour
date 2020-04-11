// Sample program to show how variables of an unnamed type can be assigned to variables of a named type,
// when they are identical.
package main

import "fmt"

// example represents a type with different fields.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// Declare a variable of an anonymous type and init using struct literal.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	// Create a value of type example.
	var ex example

	// Assign the value of the unnamed struct type to the named struct type value.
	ex = e

	// Display the values.
	fmt.Printf("%+v\n", ex)
	fmt.Printf("%+v\n", e)
	fmt.Println("Flag\t:", e.flag)
	fmt.Println("Counter\t:", e.counter)
	fmt.Println("Pi\t:", e.pi)
}
