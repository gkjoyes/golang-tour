// Sample program to show how to declare constants and their implementation in Go.
package main

func main() {

	// Untyped constants.
	const ui = 12345   // kind: integer
	const uf = 3.14159 // kind: floating-point

	// Typed constants still use the constants type system but their precision is restricted.
	const ti int = 12345       // type: int
	const tf float64 = 3.14159 // type: float64

	// Kind promotion is supported for untyped constants.

	// Constant third will be of kind floating point.
	const third = 1 / 3.0 // KindFloat(1) /  KindFloat(3.0)

	// Constant zero will be of kind integer.
	const zero = 1 / 3 // KindFloat(1) / KindFloat(3)

	// This is an example of constant arithemetic between typed and untyped constants.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
}
