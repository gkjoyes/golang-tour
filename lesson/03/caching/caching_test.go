// Tests to show how data oriented design matters.

// go test -run none -bench . -benchtime 3s
package caching

import "testing"

var count int

// Capture the time it takes to perfrom a linked list traversal.
func BenchmarkLinkedListTraverse(b *testing.B) {
	var c int

	for i := 0; i < b.N; i++ {
		c = LinkedListTraverse()
	}
	count = c
}

// Capture the time it takes to perform a row traversal.
func BenchmarkRowTraverse(b *testing.B) {
	var c int

	for i := 0; i < b.N; i++ {
		c = RowTraverse()
	}
	count = c
}

// Capture the time it takes to perfrom a column traversal.
func BenchmarkColumnTraverse(b *testing.B) {
	var c int

	for i := 0; i < b.N; i++ {
		c = ColumnTraverse()
	}
	count = c
}
