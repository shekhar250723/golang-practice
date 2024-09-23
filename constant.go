package main

import "fmt"

//Global constant variable
const Gravity = 9.8

func main() {
	//Local constant variable
	const c = "Golang"
	fmt.Println("Hey", c)
	fmt.Println("Happy", Gravity, "Day")
}
