package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50      // const never change the value
var confereneceName = "Go Conference" // : make it simpler
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // creating slice with maps

type UserData struct {
	FirstName       string
	LastName        string
	email           string
	NumberOfTickets uint
}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application\n", confereneceName) // %v type the vaule , better use %s declare the type yourself for faster processing
	fmt.Printf("We have total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	FirstNames := []string{}
	for _, booking := range bookings { // range for diff structures(not only array and slice) // _ used to identify unused var
		FirstNames = append(FirstNames, booking.FirstName)
	}
	return FirstNames
} // first () for inputs, second () for output types in order

func getUserInput() (string, string, string, uint) {
	var FirstName string
	var LastName string
	var email string
	var userTickets uint // uint not neg-
	// ask user for their informations
	fmt.Println("Enter your first name: ")
	fmt.Scan(&FirstName) // & = pointer // scan search for the reference in memory

	fmt.Println("Enter your last name: ")
	fmt.Scan(&LastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return FirstName, LastName, email, userTickets
}

func bookTicket(userTickets uint, FirstName string, LastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		FirstName:       FirstName,
		LastName:        LastName,
		email:           email,
		NumberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) // this is a slice
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s \n", FirstName, LastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %s\n", remainingTickets, confereneceName)
}

func sendTicket(userTickets uint, FirstName string, LastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, FirstName, LastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################")
}

func main() {

	greetUsers()

	for {

		FirstName, LastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(FirstName, LastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, FirstName, LastName, email)
			sendTicket(userTickets, FirstName, LastName, email)

			FirstNames := getFirstNames() // call function and save returns in var
			fmt.Printf("The first names of bookings are: %v\n", FirstNames)

			if remainingTickets == 0 {
				fmt.Println("Our Conference is booked out. Come next year.")
				break // end program
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invaild")
			}
		}
	}

}
