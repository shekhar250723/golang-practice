package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
	phone int
}

func main() {
	p := Person{"Kim", 30, "kim@test.com", 9837273929}
	fmt.Println("Name : ", p.name)
	fmt.Println("Age : ", p.age)
	fmt.Println("Email : ", p.email)
	fmt.Println("Phone : ", p.phone)
}
