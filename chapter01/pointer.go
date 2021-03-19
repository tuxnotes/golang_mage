package main

import "fmt"

func main() {
	a := 2

	var p *int = &a
	*p = 3
	fmt.Printf("%T %p %d\n", p, p, a)
}
