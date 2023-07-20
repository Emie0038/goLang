package main

import (
	"fmt"
	"strings"
)

func main() {

	// Declare variables and assign values using short declaration syntax (type inference)
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	//var bookings []string //string{"Emie", "Gloria", "Anjelie", "Jojie", "Simona", "Wanesa"}
	bookings := []string{}

	fmt.Printf("Hello, World!") // prints "Hello, world!" to the console
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		// ask user for their info
		var firstName string
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

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets, and receipts will be send to your email %v.\n", firstName, lastName, userTickets, emailAddress)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry we don't have enough tickets left at this time!")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you cannot book %v tickets.\n", remainingTickets, userTickets)
		}

	}

}
