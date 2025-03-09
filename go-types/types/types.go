package types

import "fmt"

type Animal interface {
	Sound() string
	NumberOfLegs() int
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Sound(),
		"and has", a.NumberOfLegs(), "legs")
}

func (d Dog) Sound() string {
	return "Woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}
