package main

import "fmt"

func main() {
	//Declare and Initialize an array
	var arr1 = [6]int{8, 4, 3, 2, 9, 11}
	for i := 0; i < 6; i++ {
		fmt.Print(arr1[i], " ")
	}
}
