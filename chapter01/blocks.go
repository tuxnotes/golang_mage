package main

import "fmt"

func main() {
	outer := 1
	{
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
		{
			inner := 3
			fmt.Println(inner)
		}
	}
	// 下面是不能使用的
	// fmt.Println(inner)
}
