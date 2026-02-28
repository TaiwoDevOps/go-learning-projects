package main

import (
	"fmt"
	"log"
	"os"
)

// what is the difference between dir and file path in go
// when working with files (reading from or writing to a file) always use the clean method on the path itself and then compare that path with what you've

func main() {
	filePath := "../paths-directory/1.txt"
	data := "Welcome to the Go programming."

	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File created successfully. Content :")
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = fmt.Fprintf(file, `
	------ Updated Content ------
	- Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
	- Go is expressive, concise, clean, and efficient.
	- Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines.
	- Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection.
	- It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.
	- Go has a robust standard library that makes it easy to build reliable and efficient software.
	- Go is used by many large companies such as Google, Uber, Twitch, Dropbox, and many more.
	- %s
	`, data)
	fmt.Println("File updated successfully. New Content :")
	file.Close()

	content, err = os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	// filePath2 := "/paths-directory/2.txt"

}
