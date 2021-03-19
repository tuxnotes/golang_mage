package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		/*
			每次调用input.Scan(),就会读取下一行，并将行尾的换行符去除，使用input.Text()获取内容.
			如果没有可读取的行了，则Scan()函数返回false
		*/
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Dup line: %s count: %d\n", line, n)
		}
	}
}
