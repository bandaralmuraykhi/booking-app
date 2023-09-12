package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50      // Total number of available conference tickets (constant)
var confereneceName = "Go Conference" // Name of the conference (variable)
var remainingTickets uint = 50        // Number of remaining available tickets (variable)
var bookings = make([]UserData, 0)    // Slice to store user booking data (initialized as an empty slice)

type UserData struct {
	FirstName       string // User's first name
	LastName        string // User's last name
	email           string // User's email address (note: should be "Email" to follow Go conventions)
	NumberOfTickets uint   // Number of tickets the user wants to book
}

func greetUsers() {
	fmt.Printf("Welcome to %s booking application\n", confereneceName)                                             // Display a welcome message with the conference name
	fmt.Printf("We have a total of %d tickets and %d are still available.\n", conferenceTickets, remainingTickets) // Display ticket availability
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	FirstNames := []string{}
	for _, booking := range bookings { // Iterate through user bookings
		FirstNames = append(FirstNames, booking.FirstName) // Collect user first names
	}
	return FirstNames
} // Function to retrieve and return a list of user first names

func getUserInput() (string, string, string, uint) {
	var FirstName string
	var LastName string
	var email string
	var userTickets uint // Number of tickets the user wants to book (unsigned integer)

	// Ask the user for their information
	fmt.Println("Enter your first name: ")
	fmt.Scan(&FirstName) // Read and store user's first name using fmt.Scan()

	fmt.Println("Enter your last name: ")
	fmt.Scan(&LastName) // Read and store user's last name using fmt.Scan()

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email) // Read and store user's email address using fmt.Scan()

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets) // Read and store the number of tickets using fmt.Scan()

	return FirstName, LastName, email, userTickets
}

func bookTicket(userTickets uint, FirstName string, LastName string, email string) {
	remainingTickets = remainingTickets - userTickets // Update the remaining ticket count

	var userData = UserData{
		FirstName:       FirstName, // Set user data
		LastName:        LastName,
		email:           email,
		NumberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)            // Add the user data to the bookings slice
	fmt.Printf("List of bookings is %v\n", bookings) // Display the updated list of bookings

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s \n", FirstName, LastName, userTickets, email) // Display a booking confirmation message
	fmt.Printf("%d tickets remaining for %s\n", remainingTickets, confereneceName)                                                                // Display the remaining ticket count
}

func sendTicket(userTickets uint, FirstName string, LastName string, email string) {
	time.Sleep(10 * time.Second) // Simulate a 10-second delay

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, FirstName, LastName) // Create a ticket message
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email) // Display a ticket message and email address
	fmt.Println("################")
}

func main() {

	greetUsers() // Call the greetUsers function to display a welcome message

	for {

		FirstName, LastName, email, userTickets := getUserInput() // Get user input

		// Call the ValidateUserInput function (not provided in this code) to validate input
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(FirstName, LastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, FirstName, LastName, email) // Book the ticket
			sendTicket(userTickets, FirstName, LastName, email) // Send a ticket confirmation email

			FirstNames := getFirstNames()                                   // Get a list of first names from bookings
			fmt.Printf("The first names of bookings are: %v\n", FirstNames) // Display the list of first names

			if remainingTickets == 0 {
				fmt.Println("Our Conference is fully booked. Come next year.")
				break // Exit the program if all tickets are booked
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain the '@' sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}
		}
	}

}
