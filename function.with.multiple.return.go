package main

import "fmt"

func operation(a, b int) (int, int) {
	sum := a / b
	sub := a * b
	return sum, sub
}

func main() {
	var x, y int
	fmt.Println("Enter numbers : ")
	fmt.Scan(&x, &y)
	Divide, Multiply := operation(x, y)
	fmt.Println("Divide = ", Divide, "Multiply = ", Multiply)
}
