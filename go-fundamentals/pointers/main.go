package main

import "fmt"

func main() {
	// Pointer is a variable that holds the memory address of a value
	// Note: a value here can be anything - function, variable, struct, etc.
	// Note: a pointer is a value in itself too that holds the memory address of that variable in other to reference (*pointerToVariable) the value

	count := 10

	// give the address of the value (count) to another variable
	pointerToCount := &count // "&" gives the address of the value

	// accessing the value from an address
	fmt.Println(*pointerToCount)
	fmt.Printf("the address of the pointer %d\n", pointerToCount)

	// change the value of the variable by dereferencing the address and in this case, assigning a new value
	*pointerToCount = 20
	fmt.Println(count)
}
