package main

import "fmt"

func main() {
	const name string = "kk"
	fmt.Println(name)
	// name = "aa"
	// 省略类型
	const PI = 3.14
	// 定义多个常量，类型相同
	const C1, C2 int = 1, 2
	// 定义多个常量，类型不同
	const (
		nick string = "tux"
		age  int    = 20
	)
	// 定义多个类型，省略类型
	const C3, C4 = "tux", 1
	const (
		C5 int = 1
		C6     // 与上面的保持一直
		C7
	)
	fmt.Println(C5, C6, C7)
	// 枚举，const + iota实现
	const (
		E1 int = (iota + 1) * 100
		E2
		E3
	)
	fmt.Println(E1, E2, E3)
}
