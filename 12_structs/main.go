package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age)

}

// hasBirthday method
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	}
	p.lastName = spouseLastName
}

func main() {
	person1 := Person{firstName: "Nikolay", lastName: "P", city: "Ekaterinburg", age: 30, gender: "m"}
	person2 := Person{firstName: "Alex", lastName: "Smith", city: "Pittsburg", age: 11, gender: "f"}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Petrova")

	person2.getMarried("Petrova")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
