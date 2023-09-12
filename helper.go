package main

import "strings"

func ValidateUserInput(FirstName string, LastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// Check if the first name and last name are at least 2 characters long
	isValidName := len(FirstName) >= 2 && len(LastName) >= 2
	// Check if the email contains the "@" symbol (a basic check for email format)
	isValidEmail := strings.Contains(email, "@")
	// Check if the number of user tickets is greater than 0 and less than the remaining available tickets
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets
	// Return the results of the validation checks as a tuple of boolean values
	return isValidName, isValidEmail, isValidTicketNumber

}
