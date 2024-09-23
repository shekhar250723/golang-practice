package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hello"
	fmt.Println(str1)
	str2 := "Golang Team"
	str3 := str1 + " " + str2
	fmt.Println(str3)
	greetings := []string{"Hello", "Go!"}
	fmt.Println(strings.Join(greetings, " "))
}
