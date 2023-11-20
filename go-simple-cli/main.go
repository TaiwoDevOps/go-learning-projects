package main

import "fmt"

// global variables
var totalNumOfTickets uint = 50
var eventName string = "Dot.V"

type GuestBooking struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets uint
}

func main() {
	var user = GuestBooking{}
	var bookings []GuestBooking

	fmt.Println("Welcome to the Dot.V fest!")

	fmt.Println("Please enter your first name")
	fmt.Scan(&user.firstName)
	if user.firstName == "" {
		fmt.Println("Please enter your first name")
		fmt.Scan(&user.firstName)
	}

	fmt.Println("Please enter your last name")
	fmt.Scan(&user.lastName)
	if user.lastName == "" {
		fmt.Println("Please enter your last name")
		fmt.Scan(&user.lastName)
	}

	fmt.Println("Please enter your email address")
	fmt.Scan(&user.email)
	if user.email == "" {
		fmt.Println("Please enter your email address")
		fmt.Scan(&user.email)
	}

	user.numOfTickets = takeNumOfTicket(user.numOfTickets, totalNumOfTickets)

	bookings = append(bookings, user)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", user.firstName, user.lastName, user.numOfTickets, user.email)
	fmt.Printf("%v guests list: %v\n", eventName, bookings)
}

func takeNumOfTicket(numOfTickets uint, totalNumOfTickets uint) uint {
	fmt.Println("Please enter the number of tickets")
	fmt.Scan(&numOfTickets)

	if numOfTickets > totalNumOfTickets {
		fmt.Printf("Sorry we only have %d tickets left!\n", totalNumOfTickets)
		takeNumOfTicket(numOfTickets, totalNumOfTickets)
	}
	return numOfTickets
}
