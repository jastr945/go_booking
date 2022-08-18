package main

import (
	"fmt"
)

func main() {
	confName := "Go Conference"
	const confTickets uint = 60
	var remainingTickets uint = 50

	// print types
	fmt.Printf("confName is of type %T, confTickets is of type %T\n", confName, confTickets)
	// print with new line
	fmt.Println("Welcome to the", confName, "!")
	fmt.Printf("Get your tickets here! Available tickets: %v; remaining tickets: %v\n", confTickets, remainingTickets)

	// declare a slice as the number of elements is not defined
	bookings := []string{"user2", "user3"}
	var userName string
	var userTickets int

	for {
		fmt.Println("Enter your name: ")
		// user input with a & pointer (memory location)
		fmt.Scan(&userName)

		fmt.Println("Enter the number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		bookings = append(bookings, userName)
		remainingTickets = remainingTickets - uint(userTickets)

		fmt.Printf("User %v booked %v ticket(s)\n", userName, userTickets)
		fmt.Printf("Remaining tickets after booking: %v\n", remainingTickets)
		fmt.Printf("All bookings: %v; bookings type: %T, slice length: %v\n", bookings, bookings, len(bookings))
	}
}
