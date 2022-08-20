package main

import (
	"fmt"
	"strings"
	"time"
)

const confTickets uint = 60
var confName string = "Go Conference"
var remainingTickets uint = 60
// declare a list of structs, with the initial size of zero
var bookings = make([]UserInfo, 0)

type UserInfo struct {
	user string
	ticketsBooked int
}

func main() {
	greetUser()
	for remainingTickets > 0 && len(bookings) <= int(confTickets) {
		var userName, userLastName, userTickets = getUserInput()
		var isValidInput = validateInput(userName, userLastName, userTickets, int(remainingTickets))
		if isValidInput {
			bookTickets(userName, userLastName, userTickets)
			go processEmail(userName)
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

func bookTickets(userName string, userLastName string, userTickets int) {
	var userData = UserInfo {
		user: userName + " " + userLastName,
		ticketsBooked: userTickets,
	}
	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("User %v booked %v ticket(s)\n", userName, userTickets)
	fmt.Printf("Remaining tickets after booking: %v\n", remainingTickets)
}

func printBookings() {
	// print bookings with the first name and the first letter of the last name
	for index, booking := range bookings {
		nameSplit := strings.Fields(booking.user)
		lastNameFirstLetter := nameSplit[1][:1]
		nameSplit[1] = lastNameFirstLetter + "."
		bookings[index].user = nameSplit[0] + " " + nameSplit[1]
	}
	fmt.Printf("All bookings: %v; bookings type: %T, map length: %v\n", bookings, bookings, len(bookings))
}

func processEmail(userName string) {
	// simulating a longer process
	time.Sleep(5 * time.Second)
	fmt.Printf("Email is sent to user %v\n", userName)
}
