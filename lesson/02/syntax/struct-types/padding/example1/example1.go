// Sample program to show how padding works in 32-bit architecture.
// Build: GOARCH=386 go build -o 386
// Run: ./386

package main

import (
	"fmt"
	"unsafe"
)

// Padding on 32bit Arch. Word size is 4 bytes.
type student struct {
	id     int32  // 	4 bytes				size 4
	name   string //	8 bytes				size 12
	city   string // 	8 bytes 			size 20
	passed bool   //	1 bytes				size 21
	//					3 bytes padding		size 24
	phone int32 //		4 bytes 			size 28
}

func main() {
	var s student
	size := unsafe.Sizeof(s)
	fmt.Printf("Size of[%d][%p %p %p %p %p]\n", size, &s.id, &s.name, &s.city, &s.passed, &s.phone)
}
