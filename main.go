package main

import (
	"fmt"
	"strings"
)

const confTickets uint = 60
var confName string = "Go Conference"
var remainingTickets uint = 60
// declare a slice as the number of elements is not defined
var bookings = []string{"Tom Smith", "Jane Brown"}

func main() {
	greetUser()
	for remainingTickets > 0 && len(bookings) <= int(confTickets) {
		var userName, userLastName, userTickets = getUserInput()
		var isValidInput = validateInput(userName, userLastName, userTickets, int(remainingTickets))
		if isValidInput {
			bookings = bookTickets(userName, userLastName, userTickets)
			printBookings()
			if remainingTickets == 0 {
				fmt.Println("Our event is sold out.")
				break
			}
		} else {
			fmt.Println("Please try again")
		}
	}
}

func greetUser() {
	// print types
	fmt.Printf("confName is of type %T, confTickets is of type %T\n", confName, confTickets)
	// print with new line
	fmt.Println("Welcome to the", confName, "!")
	fmt.Printf("Get your tickets here! Available tickets: %v; remaining tickets: %v\n", confTickets, remainingTickets)
}

func getUserInput() (string, string, int) {
	var userName string
	var userLastName string
	var userTickets int

	fmt.Println("Enter your first name: ")
	// user input with a & pointer (memory location)
	fmt.Scan(&userName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return userName, userLastName, userTickets
}

func validateInput(userName string, userLastName string, userTickets int, remainingTickets int) (bool) {
	var isValid bool
	isValidInput := len(userName) > 2 && len(userLastName) > 2 && userTickets > 0
	if userTickets > int(remainingTickets) && isValidInput {
		fmt.Printf("You are trying to book more tickets than available. We only have %v tickets remaining.\n", remainingTickets)
	} else if userTickets <= int(remainingTickets) && !isValidInput {
		fmt.Println("Invalid user input")
	} else {
		isValid = true
	}
	return isValid
}

func bookTickets(userName string, userLastName string, userTickets int) ([]string) {
	bookings = append(bookings, userName + " " + userLastName)
	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("User %v booked %v ticket(s)\n", userName, userTickets)
	fmt.Printf("Remaining tickets after booking: %v\n", remainingTickets)
	return bookings
}

func printBookings() {
	// print bookings with the first name and the first letter of the last name
	editedBookings := []string{}
	for _, booking := range bookings {
		nameSplit := strings.Fields(booking)
		lastNameFirstLetter := nameSplit[1][:1]
		nameSplit[1] = lastNameFirstLetter + "."
		editedBookings = append(editedBookings, nameSplit[0] + " " + nameSplit[1])
	}
	fmt.Printf("All bookings: %v; bookings type: %T, slice length: %v\n", editedBookings,  editedBookings, len(editedBookings))
}
