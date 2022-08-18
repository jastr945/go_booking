package main

import (
	"fmt"
	"strings"
)

func main() {
	confName := "Go Conference"
	const confTickets uint = 60
	var remainingTickets uint = 60

	// print types
	fmt.Printf("confName is of type %T, confTickets is of type %T\n", confName, confTickets)
	// print with new line
	fmt.Println("Welcome to the", confName, "!")
	fmt.Printf("Get your tickets here! Available tickets: %v; remaining tickets: %v\n", confTickets, remainingTickets)

	// declare a slice as the number of elements is not defined
	bookings := []string{"Tom Smith", "Jane Brown"}
	var userName string
	var userLastName string
	var userTickets int

	for remainingTickets > 0 && len(bookings) <= 50 {
		fmt.Println("Enter your first name: ")
		// user input with a & pointer (memory location)
		fmt.Scan(&userName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&userLastName)

		fmt.Println("Enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		isValidInput := len(userName) > 2 && len(userLastName) > 2 && userTickets > 0

		if userTickets > int(remainingTickets) && isValidInput {
			fmt.Printf("You are trying to book more tickets than available. We only have %v tickets remaining. Try again.\n", remainingTickets)
		} else if userTickets <= int(remainingTickets) && !isValidInput {
			fmt.Println("Invalid user input")
		}  else {
			bookings = append(bookings, userName + " " + userLastName)
			remainingTickets = remainingTickets - uint(userTickets)

			fmt.Printf("User %v booked %v ticket(s)\n", userName, userTickets)
			fmt.Printf("Remaining tickets after booking: %v\n", remainingTickets)

			// print bookings with the first name and the first letter of the last name
			editedBookings := []string{}
			for _, booking := range bookings {
				nameSplit := strings.Fields(booking)
				lastNameFirstLetter := nameSplit[1][:1]
				nameSplit[1] = lastNameFirstLetter + "."
				editedBookings = append(editedBookings, nameSplit[0] + " " + nameSplit[1])
			}

			fmt.Printf("All bookings: %v; bookings type: %T, slice length: %v\n", editedBookings,  editedBookings, len(editedBookings))

			if remainingTickets == 0 {
				fmt.Println("Our event is sold out.")
				break
			}
		}
	}
}
