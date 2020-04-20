// Sample program to show how padding works in 64-bit architecture.
// Build: GOARCH=amd64 go build -o amd64
// Run: ./amd64

package main

import (
	"fmt"
	"unsafe"
)

// Padding on 64-bit Arch. Word size is 8 bytes.
type student struct {
	id int32 // 		4 bytes				size 4
	// 					4 bytes padding		size 8
	name   string //	16 bytes			size 24
	city   string // 	16 bytes 			size 40
	passed bool   //	1 bytes				size 41
	//					3 bytes padding		size 44
	phone int32 //		4 bytes 			size 48
}

func main() {
	var s student
	size := unsafe.Sizeof(s)
	fmt.Printf("Size of[%d][%p %p %p %p %p]\n", size, &s.id, &s.name, &s.city, &s.passed, &s.phone)
}
