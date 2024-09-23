package main

import "fmt"

type Info struct {
	name       string
	age        int
	location   string
	experience bool
}

func main() {
	fmt.Println(Info{"Kim", 30, "India", false})
}
