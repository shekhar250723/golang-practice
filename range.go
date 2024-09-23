package main

import "fmt"

var pow = []int{25, 50, 80, 100, 200}

func main() {
	for i, v := range pow {
		fmt.Printf("Index = %d Value = %d\n", i, v)
	}
}
