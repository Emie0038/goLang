package main

import "fmt"

func main() {

	// Declare variables and assign values using short declaration syntax (type inference)
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string //string{"Emie", "Gloria", "Anjelie", "Jojie", "Simona", "Wanesa"}

	fmt.Printf("Hello, World!") // prints "Hello, world!" to the console
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	// ask user for their name
	var lastName string
	var emailAddress string
	var userTickets uint
	fmt.Println("Enter your firstname: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your lastname: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("Thank you %v %v for booking %v tickets, and receipts will be send to your email %v.\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
