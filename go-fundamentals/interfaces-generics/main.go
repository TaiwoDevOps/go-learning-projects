package main

import "fmt"

// interface in Go is implicit not explicit
// no "implement" keyword
// calling method just needs to satisfy the "contract"

type Person interface {
	GetName() string
}
type Numbers interface {
	int | float64 | float32 | int32 | int64
}

type Employee struct {
	ID   int
	Name string
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (e Employee) GetName() string {
	return e.Name
}

func (e BusinessPerson) GetName() string {
	return e.Name
}

func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func main() {

	displayPerson(BusinessPerson{
		ID:   2,
		Name: "Joshua Tee",
	})

	displayPerson(Employee{
		ID:   1,
		Name: "Adisa Taiwo",
	})

	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum(1.1, 2.2, 3.3, 4.4, 5.5))
}

// Generics
func Sum[T Numbers](numbers ...T) T {
	total := T(0)
	for _, num := range numbers {
		total += num
	}
	return total
}
