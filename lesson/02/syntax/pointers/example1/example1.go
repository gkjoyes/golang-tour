// Sample program to show the basic concept of pass by value.
package main

func main() {

	// Declare a variable of type int with a value of 10.
	count := 10

	// Pass the "value of" the count.
	increment(count)

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

}

// increment declares inc variable of type int and performs some modifications.
func increment(inc int) {

	// Increment the "value of" inc.
	inc++

	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}
