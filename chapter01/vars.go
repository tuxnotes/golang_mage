package main

import "fmt"

var version string = "1.0" // 函数外定义的变量可以不使用

func main() {
	// define variable named me,type string
	/*
		varible name should met requirements for identifier
		1. consist of letter(unicode),number,and underscore
		2. can't start with numbers
		3. can't use keyword
		4. avoid conflict with build-in indentifier
		5. camel case
		6. case sensitive
	*/
	var me string // 函数内定义的变量必须使用
	me = "KK"
	fmt.Println(me)
	fmt.Println(version)
	var user, name string = "KK", "woniu"
	fmt.Println(user, name)
	var (
		age    int // 注意后面没有逗号
		height float64
	)
	fmt.Println(age, height)
	// 短声明只能在函数内部使用
	isBody := true
	fmt.Println(isBody)
}
