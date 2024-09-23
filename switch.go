package main

import "fmt"

func main() {
	/* local variable definition */
	var grade string
	var marks int

	fmt.Print("Enter marks : ")
	fmt.Scanln(&marks)

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	fmt.Printf("Your grade is %s\n", grade)
}
