package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "Sharon@gmail.com"
	emails["Vasya"] = "Vasya@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	delete(emails, "Bob")
	fmt.Println(emails)

	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "Sharon@gmail.com"}
	fmt.Println(emails2)

}
