// Sample program to show how to declare and initialize struct types.
package main

import "fmt"

// example represent a type with different field.
type example struct {
	flag  bool
	count int16
	pi    float32
}

func main() {

	// Declare a variable of type example set to its zero value.
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using a struct literal.
	e2 := example{
		flag:  true,
		count: 10,
		pi:    3.14159,
	}

	// Display the field values.
	fmt.Println("Flag\t:", e2.flag)
	fmt.Println("Counter\t:", e2.count)
	fmt.Println("Pi\t:", e2.pi)
}
