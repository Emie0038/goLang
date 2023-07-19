package main

import "fmt"

func main() {

	// Declare variables and assign values using short declaration syntax (type inference)
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainningTickets = 50

	fmt.Printf("Hello, World!") // prints "Hello, world!" to the console
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainningTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	// ask user for their name
	var userTickets int

	userName = "Bill"
	userTickets = 12
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
