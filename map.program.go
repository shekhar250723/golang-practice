package main

import "fmt"

func main() {
	var products map[string]int
	products = make(map[string]int)

	products["A"] = 40
	products["B"] = 60
	products["C"] = 80
	fmt.Println("Product A Value = ", products["A"])

	for prod := range products {
		fmt.Println("Product Index = ", prod, " Product value = ", products[prod])
	}
}
