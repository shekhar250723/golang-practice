package main

import "fmt"

func main() {
	for i := 1; i < 40; i++ {
		if i < 5 {
			fmt.Println(i)
		} else {
			break
		}
	}
}
