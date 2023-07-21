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

	greetingsUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		// ask user for their info
		firstName, lastName, emailAddress, userTickets := getUserInput()

		//call function validateUser
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets, and receipts will be send to your email %v.\n", firstName, lastName, userTickets, emailAddress)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			//call function firstnames
			firstNames := getFirstNames(bookings)
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

func greetingsUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, emailAddress string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
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
