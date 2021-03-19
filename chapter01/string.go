package main

import "fmt"

func main() {
	var name string = "Jhon\tSmith"
	var desc string = `中\t国`
	var name1 string = "Jhon\\tSmith"
	fmt.Println(name)
	fmt.Println(desc)
	fmt.Println(name1)
	var desc1 string = "abcdef"
	fmt.Printf("%T %c\n", desc1[0], desc1[0])
	fmt.Printf("%T %s\n", desc1[0:2], desc1[0:2])
	var desc2 string = "我爱中国"
	fmt.Printf("%T %c\n", desc2[0], desc2[0])
	fmt.Printf("%T %s\n", desc2[0:2], desc2[0:2])
}
