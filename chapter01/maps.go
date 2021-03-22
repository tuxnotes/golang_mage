package main

import "fmt"

func main() {
	var scores map[string]int
	fmt.Printf("%T %#v\n", scores, scores)
	var ages map[string]int
	fmt.Println(ages == nil)    // true
	fmt.Println(len(ages) == 0) // true
	ages["carol"] = 21          // panic: assignment to entry in nil map
}
