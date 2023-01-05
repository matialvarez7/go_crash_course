package main

import "fmt"

func main() {
	// Define maps
	//emails := make(map[string]string)

	// Assign key and value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// declare a map an add key and value
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)        // Show all the elements of the map
	fmt.Println(emails["Bob"]) // Show the value for the Key "Bob"
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

}
