package main

import "fmt"

func main() {
	a := 5
	var c int
	c = a
	fmt.Println(c)
	c += a
	fmt.Println(c)
	c -= a
	fmt.Println(c)
	c *= a
	fmt.Println(c)
	c /= a
	fmt.Println(c)
	c %= a
	fmt.Println(c)
}
