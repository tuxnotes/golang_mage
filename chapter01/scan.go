package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入名字：")
	fmt.Scan(&name) // scan receive pointer
	fmt.Printf("你的名字是: %s\n", name)
}
