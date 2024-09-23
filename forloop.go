package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number : ")
	fmt.Scan(&n)
	for i := 20; i < n; i++ {
		fmt.Println("Loop ", i)
	}
}
