package main

import "fmt"

type Address struct {
	City string
}

type GenericUser interface {
	disable()
}

type User struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

type Company struct {
}

func (c Company) disable() {
	fmt.Println("Company is disabled")
}

func (u User) disable() {
	u.Active = false
	fmt.Printf("User %v has been disabled.\n", u)
}

type ComplexAddress struct {
	Name   string
	Number int
	Address
}

func main() {

	newUser := User{
		Name:   "Junior",
		Age:    10,
		Active: true,
		Address: Address{
			City: "Maringá",
		},
	}

	fmt.Printf("Name: %s\n", newUser.Name)
	fmt.Printf("Age: %d\n", newUser.Age)
	fmt.Printf("Active: %t\n", newUser.Active)
	fmt.Printf("Address: %s\n", newUser.Address.City)
	newUser.disable()

	complexAddress := ComplexAddress{}
	complexAddress.Name = "Test"
	complexAddress.Number = 100
	complexAddress.City = "Maringa"

	fmt.Printf("Name: %s\n", complexAddress.Name)
	fmt.Printf("Number: %d\n", complexAddress.Number)
	fmt.Printf("Address: %s\n", complexAddress.City)

	company := Company{}
	company.disable()

}
