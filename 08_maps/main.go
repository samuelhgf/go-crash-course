package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)

	// Assign key values
	//emails["Bob"] = "bob@gmail.com"
	//emails["Samuel"] = "samuel@gmail.com"
	//emails["Sharon"] = "sharon@gmail.com"

	// Declare map and key values
	emails := map[string]string{"Bob":"bob@gmail.com", "Samuel":"samuel@gmail.com", "Sharon":"sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")

	fmt.Println(emails)
}
