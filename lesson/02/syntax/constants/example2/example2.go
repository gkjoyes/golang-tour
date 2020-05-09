// Sample program to show how constants do have at least 256 bits of precision.
package main

import "fmt"

const (
	// Max integer value on 64 bit architecture.
	maxInt = 9223372036854775807

	// Much larger value than int64.
	bigger = 92233720368547758072356978

	// Will NOT	compile.
	// biggerInt int64 = 92233720368547758072356978
)

func main() {
	fmt.Println("Will Compile")
}
