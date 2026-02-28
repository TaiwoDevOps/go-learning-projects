package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func (e *Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) SetStatus(active bool) {
	e.IsActive = active
}

func NewEmployee(id int, firstName string, lastName string, position string, salary int, isActive bool, joinedAt time.Time) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  joinedAt,
	}
}

func main() {
	user1 := NewEmployee(1, "Jane", "Doe", "Developer", 50000, true, time.Now())
	println(user1.FullName())
	user1.FirstName = "Taiwo"
	user1.LastName = "Adisa"
	println(user1.FullName())
	user1.SetStatus(false)
	println(user1.IsActive)

	println("=================Example 2 ================")
	user2 := Employee{
		ID:        2,
		FirstName: "John",
		LastName:  "Doe",
		Position:  "Manager",
		Salary:    65000,
		IsActive:  true,
		JoinedAt:  time.Now().Add(time.Hour * 72),
	}

	println(user2.FullName())
	user2.SetStatus(false)
	println(user2.IsActive)
	fmt.Printf("User 2 is : %+v", user2)
}
