package  main

import "fmt"

func main() {
	var confName = "Go Conference"
	const confTickets = 60
	var remainingTickets = 50

	fmt.Println("Welcome to the", confName, "!")
	fmt.Printf("Get your tickets here! Available tickets: %v; remaining tickets: %v\n", confTickets, remainingTickets)
}
