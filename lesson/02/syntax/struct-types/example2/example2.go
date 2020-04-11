// Sample program to show how to declare and initialize anonymous struct types.
package main

import "fmt"

func main() {

	// Declare a variable of an anonymous type set to its zero value.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init using a struct literal.
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	// Display the values.
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag\t:", e2.flag)
	fmt.Println("Counter\t:", e2.counter)
	fmt.Println("Pi\t:", e2.pi)
}
