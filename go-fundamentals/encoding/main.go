package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	Password string `json:"-" xml:"-"`
	IsActive bool   `json:"active" xml:"active"`
}

var payload = `{"name":"John Smith","age":42,"phone":"","active":true}`

func main() {
	u := user{
		Name:  "John Smith",
		Age:   45,
		Phone: "13812345678",
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&u); err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(strings.NewReader(payload))
	if err := dec.Decode(&u); err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)

}
