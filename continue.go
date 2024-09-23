package main

import "fmt"

func main() {
	for i := 2; i < 50; i++ {
		if i < 20 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
