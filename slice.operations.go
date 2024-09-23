package main

import "fmt"

func main() {
	s := []int{20, 33, 55, 766, 112, 137}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
