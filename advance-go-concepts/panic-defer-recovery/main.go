package main

import "fmt"

func shouldPanic() {
	panic("Something just happen right now....")
}

func recoverable() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered from panic: %v\n", err)
		}
	}()
	shouldPanic()
}

func main() {

	recoverable()
}
