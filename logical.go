package main

import "fmt"

func main() {
	a, b, c, d := 20, 30, 45, false
	fmt.Println((a > b) && (b > c))
	fmt.Println((a > b) || (b > c))
	fmt.Println(!d)

	num := 25
	fmt.Println("Number :", num)
	fmt.Printf("Number : %v\n", num)
}
