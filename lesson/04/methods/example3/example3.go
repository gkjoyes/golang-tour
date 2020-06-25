// Sample program to show how to declare function variables.
package main

import "fmt"

// data is struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d data) displayName() {
	fmt.Println("My name is ", d.name)
}

// setAge sets the age and display the value.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println("Name: ", d.name, "Age: ", d.age)
}

func main() {

	// Declare a variable of type data.
	d := data{
		name: "george",
	}

	fmt.Println("Proper calls to methods")
	call(d)
	mechanism(d)

	fmt.Println("\nCall value receiver method with variable:")
	functionVariable1(d)

	fmt.Println("\nCall Pointer receiver method with variable:")
	functionVariable2(d)
}

// How we actually call methods in Go.
func call(d data) {

	d.displayName()
	d.setAge(25)
}

// This is what Go is doing underneath (how things work).
func mechanism(d data) {

	data.displayName(d)
	(*data).setAge(&d, 23)
}

func functionVariable1(d data) {

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method is using a value receiver.
	f := d.displayName

	// Call the method via the variable.
	f()

	// Change the value of d.
	d.name = "tony"

	// Call the method via the variable. We don't see the change.
	f()
}

func functionVariable2(d data) {

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method is using a pointer receiver.
	f := d.setAge

	// Call the method via the variable.
	f(25)

	// Change the value of d.
	d.name = "manu"

	// Call the method via the variable. We see the change.
	f(25)
}
