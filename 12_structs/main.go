package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Greeting method (value receiver)
func (p Person) Greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// HasBirthday method (pointer receiver)
func (p *Person) HasBirthday() {
	p.age++
}

// GetMarried (pointer receiver)
func (p *Person) GetMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	}else {
		p.lastName += " " + spouseLastName
	}
}

func main() {
	// Init Person using struct
	person1 := Person{firstName: "Samuel", lastName: "França", city: "São José dos Campos", gender: "M", age: 21}
	person2 := Person{firstName: "Thais", lastName: "Barros", city: "Campina Grande", gender: "F", age: 19}

	// Alternative
	//person2 := Person{"Samuel", "França", "São José dos Campos", "M", 21}
	//
	//fmt.Println(person1)
	//fmt.Println(person2)
	//fmt.Println(person2.age)

	person1.HasBirthday()
	person2.GetMarried("França")
	fmt.Println(person1.Greet())
	fmt.Println(person2.Greet())
}
