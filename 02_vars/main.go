package main

import "fmt"


func main() {
	// Using var
	//var name = "Brad"
	var age = 21
	const isCool = true // can't modify

	// Shorthand
	// name := "Brad"
	// email := "brad@gmail.com"
	size := 1.3

	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)

	// Output the type of the variable
	fmt.Printf("%T\n", name) // string
	fmt.Printf("%T\n", age) // int
	fmt.Printf("%T\n", isCool) // bool
	fmt.Printf("%T\n", size) // float64
}
