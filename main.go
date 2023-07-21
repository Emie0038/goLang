package main

import (
	"fmt"
	"strings"
)

// Declare variables and assign values using short declaration syntax (type inference)
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

//var bookings []string //string{"Emie", "Gloria", "Anjelie", "Jojie", "Simona", "Wanesa"}

func main() {

	greetUsers()

	for {
		// ask user for their info
		firstName, lastName, emailAddress, userTickets := getUserInput()

		//call function validateUser
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailAddress, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//call booking function
			bookedTickets(userTickets, firstName, lastName, emailAddress)

			//call function firstnames
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry we don't have enough tickets left at this time!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name must be longer than two characters.")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email adress with '@' symbol in it.")
			}
			if !isValidTicketNumber {
				fmt.Printf("You can only buy between one and %v tickets.\n", conferenceTickets)
			}
			fmt.Printf("We have %v tickets remaining. Try again!.\n", remainingTickets)
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, emailAddress, userTickets
}

func bookedTickets(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, and receipts will be send to your email %v.\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
