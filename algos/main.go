package main

import "fmt"

func main() {
	var name string
	var age int
	var accountBalance float64
	var married bool
	var pHolder string

	name = "Jalo Tems"
	age = 25
	accountBalance = 120000000
	married = false

	if married {
		pHolder = "I'm"
	} else {
		pHolder = "Not"
	}

	fmt.Printf("My name is: %s\nI am %d old\n%s married\nBut my current balance is: %.2f", name, age, pHolder, accountBalance)

}
