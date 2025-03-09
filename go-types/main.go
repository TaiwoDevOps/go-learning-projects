package main

import 
(
	"github.com/TaiwoDevOps/go-types/types"
	"fmt"
)



// see if you can serialize and deserialize a struct to json

func main() {
	// create a list of user
	users := []User{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "j@j.com",
			Age:       30,
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "j@j.com",
			Age:       30,
		},
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "j@j.com",
			Age:       30,
		},
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "j@j.com",
			Age:       30,
		},
	}

	fmt.Println(users)

	dog := Dog{
		Name:  "Rico",
		Breed: "Lahsa Alpso",
	}

	PrintInfo(dog)

}

