package main

import (
	"strings"
)

func validateUserInput(firstUser string, lastUser string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstUser) >= 2 && len(lastUser) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets

}
