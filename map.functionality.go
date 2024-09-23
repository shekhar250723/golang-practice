package main

import "fmt"

func main() {
	sub := make(map[string]int)

	sub["math"] = 45
	fmt.Println("The value:", sub["math"])

	sub["math"] = 65
	fmt.Println("The value:", sub["math"])

	delete(sub, "math")
	fmt.Println("The value:", sub["math"])

	v, ok := sub["math"]
	fmt.Println("The value:", v, "Present?", ok)
}
