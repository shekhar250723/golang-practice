package main

import "fmt"

func main() {
	primes := [8]int{7, 4, 66, 33, 23, 1, 456, 945}

	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println(primes[:4])
	fmt.Println(primes[3:])
}
