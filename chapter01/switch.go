package main

import "fmt"

func conflip() string {
	return "heads"
}

func Signum(x int) int {
	switch { // tagless switch, which is equivalent to switch true
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func main() {
	var heads, tails int
	switch conflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}
}
