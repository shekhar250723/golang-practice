package main

import "fmt"

func sqr(x int) {
	result := x + x
	fmt.Print("Addition : ", result)
}

func main() {
	var num int
	fmt.Print("Enter number : ")
	fmt.Scan(&num)
	sqr(num)
}
