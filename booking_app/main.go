package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("confrenceTicket is %T \nremainingTickets is %T\nconferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets uint

	userName = "John Doe"

	println(userName)
	println(&userName)
	
	fmt.Scan(&userName)
	println(userName)

	fmt.Scan(&userTickets)
	fmt.Printf("user %v booked %v tickets", userName, userTickets)
}