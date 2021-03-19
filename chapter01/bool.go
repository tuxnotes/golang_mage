package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var zero bool
	isBoby := true
	fmt.Printf("Type of isBoby: %T, bytes take for bool type: %d\n", isBoby, unsafe.Sizeof(isBoby))
	fmt.Println(zero)
	fmt.Printf("Default zero value for bool type: %v\n", zero)
}
