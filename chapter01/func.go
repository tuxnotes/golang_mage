package main

import "fmt"

func SayHello() {
	fmt.Println("hello world")
}

func main() {
	fmt.Printf("%T\n", SayHello)
	SayHello()
}
