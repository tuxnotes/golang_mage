package main

import "fmt"

func main() {
	var numbers [10]int
	fmt.Printf("%T\n", numbers)
	fmt.Println(numbers)
	var s1 [3]string
	fmt.Printf("%q\n", s1)
	nums1 := [...]int{1, 2, 3}
	nums2 := [...]int{1, 2, 4}
	fmt.Println(nums1 == nums2)
}
