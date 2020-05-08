// Sample program to show the basic concept of using a pointer to share data.

package main

func main() {

	// Declare a variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")

	// Pass the "address of" count.
	increment(&count)

	println("count:\tValue of[", count, "]\tAddr of[", &count, "]")
}

// increment declares inc as a pointer variable of type int whose value is always an address.
func increment(inc *int) {

	// Increment the "value of" inc that the "pointer points to".
	*inc++

	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
