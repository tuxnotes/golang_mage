package main

import "fmt"

func main() {
	var b []byte = []byte{'a', 'b', 'c'}
	fmt.Printf("%T %v\n", b, b)
	s := string(b)
	fmt.Printf("%T %v\n", s, s)
	bs := []byte(s)
	fmt.Printf("%T %v\n", bs, bs)
}
