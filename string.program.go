package main

import "fmt"

func main() {
	var str string = "Hello Golang Team"
	fmt.Println(str)
	fmt.Printf("%s\n", str)

	fmt.Println("Length : ", len(str))
}
