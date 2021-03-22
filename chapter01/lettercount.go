package main

import "fmt"

func main() {
	article := `
	hello world
	how are you
	...
	`
	articleStat := map[rune]int{}
	for _, ch := range article {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			articleStat[ch]++
		}
	}
	for c, n := range articleStat {
		fmt.Printf("%c : %d\n", c, n)
	}
}
