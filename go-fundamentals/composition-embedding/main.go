package main

import (
	"fmt"
	"strings"
)

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if strings.Trim(a.Street, " ") == "" && strings.Trim(a.City, " ") == "" {
		return "No address provided"
	}

	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type ContactInfo struct {
	Email string
	Phone string
}

func (c ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", c.Email, c.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Name: %s\nAddress: %s\nContact: %s\nBusiness Type: %s\n", c.Name, c.Address.FullAddress(), c.ContactInfo.DisplayContact(), c.BusinessType)
}

type CompanyWithPersonalEmail struct {
	Name string
	Address
	ContactInfo
	PersonalEmail string
}

func main() {
	fmt.Println("-------- Struct Embedding -----------")
	comp := Company{
		Name: "Company Name",
		Address: Address{
			Street:  "   ",
			City:    "   ",
			State:   "Company State",
			ZipCode: "Company ZipCode",
		},
		ContactInfo: ContactInfo{
			Email: "Company Email",
			Phone: "Company Phone",
		},
		BusinessType: "Company Business Type",
	}

	comp.GetProfile()
}
