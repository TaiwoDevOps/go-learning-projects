package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextId int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(c *Contact) {
	if _, exist := contactIndexByName[c.Name]; exist {
		fmt.Println("Contact already exist")
		return
	}
	contactList = append(contactList, *c)
	contactIndexByName[strings.ToLower(c.Name)] = len(contactList) - 1
	nextId++
}

// strings.Contains(strings.ToLower(c.Name), n) ||
func findContactByName(name string) ([]*Contact, error) {
	lN := strings.ToLower(name)
	foundContacts := make([]*Contact, 0)
	for k, v := range contactIndexByName {
		if strings.Contains(strings.ToLower(k), lN) {
			fmt.Println("Contact found at index", v)
			foundContacts = append(foundContacts, &contactList[v])
		}
	}
	if len(foundContacts) == 0 {
		return nil, fmt.Errorf("")
	}
	return foundContacts, nil
}

func listContacts() {
	fmt.Println("=================| My Contact List |=================")
	if len(contactList) == 0 {
		fmt.Println("No contacts found")
		return
	}
	for i, c := range contactList {
		fmt.Printf("%d.\nID: %d\nName: %s\nEmail: %s\nPhone: %s\n\n", i+1, c.ID, c.Name, c.Email, c.Phone)
	}
	fmt.Println("")
}

func main() {
	fmt.Println("==============| My Contact List Program |=================")
	fmt.Printf("The contact:-\nlist: %d\ncapacity: %d\n ", len(contactList), cap(contactList))

	addContact(&Contact{
		ID:    nextId,
		Name:  "Ravi",
		Email: "a@b.com",
		Phone: "1234567890",
	})
	addContact(&Contact{
		ID:    nextId,
		Name:  "Joshua",
		Email: "a@b.com",
		Phone: "1234567890",
	})
	addContact(&Contact{
		ID:    nextId,
		Name:  "Adams",
		Email: "a@b.com",
		Phone: "1234567890",
	})
	addContact(&Contact{
		ID:    nextId,
		Name:  "Taiwo Adisa",
		Email: "a@b.com",
		Phone: "1234567890",
	})
	addContact(&Contact{
		ID:    nextId,
		Name:  "Taiwo",
		Email: "a@b.com",
		Phone: "1234567890",
	})

	listContacts()

	var contactName = "taiwo-"

	contacts, err := findContactByName(contactName)
	if err != nil {
		fmt.Printf("Contact not found for: %s with count: %d and capacity as %d\n", contactName, len(contacts), cap(contacts))
	} else {
		fmt.Printf("Contact found for: %s with count: %d\n", contactName, len(contacts))
		for _, contact := range contacts {
			fmt.Printf("ID: %d\nName: %s\nEmail: %s\nPhone: %s\n\n", contact.ID, contact.Name, contact.Email, contact.Phone)
		}
	}
}
