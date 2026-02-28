package main

import "fmt"

func main() {
	fmt.Println("hello from switch")

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		default:
			fmt.Printf("Unknown type\n")
		}
	}

	checkType(5)
	checkType("hello")
	checkType(true)
}
