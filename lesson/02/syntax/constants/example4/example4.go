// Sample program to show how literal, constant and variable work within the scope of implicit conversion.
package main

import (
	"fmt"
	"time"
)

func main() {

	// Use the time package to get the current date/time.
	now := time.Now()

	// Substract 5 nanoseconds from now using a literal constant.
	literal := now.Add(-5)

	// Substract 5 seconds from now using a declared constant.
	const timeout = 5 * time.Second // time.Duration(5) * time.Duration(1000000000)
	constant := now.Add(-timeout)

	// Substract 5 nanoseconds from now using a variable of type int64.
	// Error: cannot use minusFive (type int64) as type time.Duration in argument to now.Add
	// minusFive := int64(-5)
	// variable := now.Add(minusFive)

	// Display the values.
	fmt.Printf("Now		:%v\n", now)
	fmt.Printf("Literal		:%v\n", literal)
	fmt.Printf("Constant	:%v\n", constant)
}
