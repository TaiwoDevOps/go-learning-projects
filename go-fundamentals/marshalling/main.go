package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string  `json:"fullName" xml:"fullname"`
	Password string  `json:"-" xml:"password"`
	Age      int     `json:"age" xml:"age"`
	Phone    string  `json:"phoneNumber" xml:"phone_number"`
	IsActive bool    `json:"active" xml:"is_active"`
	Profile  profile `json:"profile" xml:"profile"`
}
type profile struct {
	URL string `json:"url" xml:"url"`
}

func main() {

	user1 := user{
		Name:     "Jane",
		Age:      25,
		IsActive: false,
		Phone:    "+234 812 243 7265",
		Password: "P@ssw0rd",
		Profile: profile{
			URL: "https://www.newwebsite.com",
		},
	}

	byteSlice, err := json.Marshal(user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteSlice))

	var u user
	err = json.Unmarshal(byteSlice, &u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the user : %+v\n", u)

}
