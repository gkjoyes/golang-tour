// Sample program to show how to takes slices of slices to create different views of and make
// changes to the underlying array.
package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements and a capacity of 8.
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	inspectSlice(slice1)

	// Get slice of slice.
	slice2 := getSliceOfSlice(slice1, 2, 3)
	inspectSlice(slice1)
	inspectSlice(slice2)

	// Get slice of slice by limiting the index boundary.
	slice3 := getSliceOfSliceWithIndex(slice1, 2, 3, 4)
	inspectSlice(slice3)

	// Get copy of the slice.
	slice4 := getCopyOfSlice(slice1)
	inspectSlice(slice4)
}

// getSliceOfSlice returns slice of the slice.
func getSliceOfSlice(slice1 []string, start, end int) []string {

	// Parameters are[staring_index : (starting_index + index)].
	slice2 := slice1[start : end+1]

	// Change the value of index 0 of slice2.
	slice2[0] = "CHANGED"

	return slice2
}

// getSliceOfSliceWithIndex returns slice of the slice by limiting index boundary.
func getSliceOfSliceWithIndex(slice1 []string, start, end, index int) []string {

	// The "index" should be greater than the "end".
	return slice1[start:end:index]
}

// getCopyOfSlice make a new slice big enough to hold elements of slice1 and copy the values
// using the built-in copy function.
func getCopyOfSlice(slice1 []string) []string {

	slice2 := make([]string, len(slice1))
	copy(slice2, slice1)
	return slice2
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("\nLength[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
}
