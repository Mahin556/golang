package main

import (
	"fmt"
	"strings"
)


func main() {
	str := "  Go  is a  powerful language "
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var userTickets uint
	var userEmail string
	// var bookings [50]string //array
	var bookings []string //slice
	// Other ways to declare slice
	// var bookings1 = []string{}
	// var bookings2 []string
	// bookings3 := []string{}

	fmt.Println(strings.TrimSpace(str))

	fmt.Printf("confrenceTicket is %T \nremainingTickets is %T\nconferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)

	for {
		fmt.Println("Get your tickets here to attend")
			
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
			
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
			
		fmt.Println("Enter your email:")
		fmt.Scan(&userEmail)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)
		
		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)
			
			fmt.Printf("The Whole array %v\n", bookings)
			fmt.Printf("The first value %v\n", bookings[0])
			fmt.Printf("The array length is %v\n", len(bookings))
			fmt.Printf("The array type is %T\n", bookings)
			
			fmt.Printf("User %v %v booked %v tickets, The confirmation email will be sent to %v\n", firstName, lastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v confrence\n", remainingTickets, conferenceName)
			
			firstNames := []string{}
	
			for _, booking := range bookings {
				fields := strings.Fields(booking)
				firstNames = append(firstNames, fields[0])
			}
	
			fmt.Println("All bookings: ", firstNames)
	
			// noTicketsRemaining := remainingTickets == 0
	
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %d tickets remaining, so you can't book %d tickets\n", remainingTickets, userTickets)
		}
		
	}
}