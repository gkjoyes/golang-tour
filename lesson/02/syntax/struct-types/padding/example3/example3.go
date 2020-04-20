// Sample program to show how padding works when we don't specify ARCH during a build.
package main

import (
	"fmt"
	"unsafe"
)

// Three byte padding.
type student struct {
	id     int32 // 	4 bytes				size 4
	passed bool  //		1 bytes				size 5
	// 					3 bytes padding 	size 8
	phone int32 //		4 bytes 			size 12
}

func main() {
	var s student
	size := unsafe.Sizeof(s)
	fmt.Printf("Size of[%d][%p %p %p]\n", size, &s.id, &s.passed, &s.phone)
}
