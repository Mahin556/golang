package main

import (
	"fmt"
	"strings"
)

// Constants
const conferenceName = "Go Conference"
const conferenceTickets = 50

// Struct for booking
type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {
	remainingTickets := uint(conferenceTickets)
	bookings := make([]UserData, 0)

	greetUsers(remainingTickets)

	for remainingTickets > 0 {
		fmt.Println("\nGet your tickets here to attend")

		userData := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber :=
			validateUserInput(userData, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets -= userData.userTickets
			bookings = append(bookings, userData)

			fmt.Printf(
				"\nUser %v %v booked %v tickets. Confirmation will be sent to %v\n",
				userData.firstName,
				userData.lastName,
				userData.userTickets,
				userData.email,
			)

			fmt.Printf("%v tickets remaining for %v\n",
				remainingTickets,
				conferenceName,
			)

			firstNames := getFirstNames(bookings)
			fmt.Println("Bookings:", firstNames)

			if remainingTickets == 0 {
				fmt.Println("\nOur conference is booked out. Come back next year.")
				break
			}

		} else {
			printValidationErrors(isValidName, isValidEmail, isValidTicketNumber, remainingTickets, userData.userTickets)
		}
	}
}

// Greeting function
func greetUsers(remainingTickets uint) {
	fmt.Printf("Welcome to the %s booking application\n", conferenceName)
	fmt.Printf("We have %d tickets and %d are available\n",
		conferenceTickets,
		remainingTickets,
	)
}

// Get user input
func getUserInput() UserData {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return UserData{
		firstName:   strings.TrimSpace(firstName),
		lastName:    strings.TrimSpace(lastName),
		email:       strings.TrimSpace(email),
		userTickets: userTickets,
	}
}

// Validate input
func validateUserInput(user UserData, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(user.firstName) >= 2 && len(user.lastName) >= 2
	isValidEmail := strings.Contains(user.email, "@")
	isValidTicketNumber := user.userTickets > 0 && user.userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

// Print errors
func printValidationErrors(isValidName bool, isValidEmail bool, isValidTicketNumber bool, remainingTickets uint, userTickets uint) {

	if !isValidName {
		fmt.Println("First name or last name is too short")
	}
	if !isValidEmail {
		fmt.Println("Email must contain @")
	}
	if !isValidTicketNumber {
		fmt.Printf("We only have %d tickets remaining, you requested %d\n",
			remainingTickets,
			userTickets,
		)
	}
}

// Get first names safely
func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}