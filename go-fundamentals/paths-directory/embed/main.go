package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed public
var public embed.FS

func main() {

	content, err := public.ReadFile("public/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

}
