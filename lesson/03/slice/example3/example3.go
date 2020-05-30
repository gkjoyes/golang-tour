// Sample program to show how to grow a slice using the built-in function append and
// how append grows the capacity of the underlying array.
package main

import "fmt"

func main() {

	// Declare a nil slice of strings.
	var data []string

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~100 strings to the slice.
	for i := 0; i < 1e5; i++ {

		// Use the built-in function append to add to the slice.
		value := fmt.Sprintf("Record: %d", i)

		data = append(data, value)

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {

			// Calculate the percent of change.
			p := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)

			// Display the result.
			fmt.Printf("[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], i, cap(data), p)
		}
	}
}
