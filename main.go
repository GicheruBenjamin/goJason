package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hi there!")
	fmt.Println("In here there is arecord of people.")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Add a new person")
	fmt.Println("2. Delete a person")
	fmt.Println("3. Update a person")
	fmt.Println("4. List all people")
	fmt.Println("5. Search for a person")
	fmt.Println("6. Exit")

	var action int
	fmt.Scan(&action)

	//check if the input is valid
	if action < 1 || action > 6 {
		fmt.Println("Invalid input")
		return
	}
	//switch as per the options
	switch (action){
	case 1:
		fmt.Println("Adding a new person")
	case 2:
		fmt.Println("Deleting a person")
	case 3:
		fmt.Println("Updating a person")
	case 4:
		fmt.Println("Listing all people")
	case 5:
		fmt.Println("Searching for a person")
	case 6:
		fmt.Println("Exiting")
		fmt.Println("Bye!")
	default:
		fmt.Println("Invalid input")
	}
}