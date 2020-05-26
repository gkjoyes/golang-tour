// Package caching provides code to show why Data Oriented Design matters.
// How data layouts matter more to performance than algorithm efficiency.
package caching

import "fmt"

const (
	rows = 4029
	cols = 4029
)

// Create a square matrix of 4096 x 4096.
var matrix [rows][cols]byte

// data represent a data node for our linked list.
type data struct {
	v byte
	p *data
}

// list points to the head of the list.
var list *data

func init() {
	var last *data

	// Create a linked list with the same number of elements.
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			// Create a new node and link it in.
			var d data
			if list == nil {
				list = &d
			}

			if last != nil {
				last.p = &d
			}
			last = &d

			// Add values to all even rows.
			if r%2 == 0 {
				matrix[r][c] = 0xFF
				d.v = 0xFF
			}
		}
	}

	// Count number of elements in the linked list.
	var c int
	d := list
	for d != nil {
		c++
		d = d.p
	}

	fmt.Println("Elements in the linked list:", c)
	fmt.Println("Elements in the matrix:", rows*cols)
}

// LinkedListTraverse traverses the linked list linearly.
func LinkedListTraverse() int {
	var count int

	d := list
	for d != nil {
		if d.v == 0xFF {
			count++
		}
		d = d.p
	}
	return count
}

// ColumnTraverse traverses the matrix linearly down each column.
func ColumnTraverse() int {
	var count int

	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			if matrix[r][c] == 0xFF {
				count++
			}
		}
	}
	return count
}

// RowTraverse traverses the matrix linearly down each row.
func RowTraverse() int {
	var count int

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0xFF {
				count++
			}
		}
	}
	return count
}
