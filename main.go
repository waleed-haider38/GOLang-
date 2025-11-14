package main

import (
	"fmt"
	"strings"
)

// fmt is package where we can use some built in functions like Print , Println etc there are multiple package which we see next in our tutorial.
// print is used to print some statement but println is also used to print statement but it starts a new line. Now we use printf function which is used to identify the varialbes in effiecient way.
// %v is used to show the value of variable but %T is used to show the type of variable.
// := is only used with var and we do not need to explicity identify the type of variable so it save our time.
// & this sign is concept of pointer use to store the address of a variable in go language. and variables are store in memory so it refers to our variable addressed where our variables or values are store.
func main() {
	myapp()
	userName := "Waleed Haider"
	// userAge := 23
	bookings := []string{}

	// different ways to declare the variables

	// var myUser string = "Go Conference"
	// const myage int = 18

	// fmt.Println(userAge)
	// fmt.Println(&userAge)
	// fmt.Printf("userName type is %T Age type is %T \n", userName, userAge)
	// fmt.Printf("I am %v years old\n", userAge)

	var remainingTickets uint = 50

	greetuser(userName, remainingTickets)

	for {

		firstUser, lastUser, email, userTickets := getUserInput()

		// calling function to validate user input
		isValidName, isValidEmail, isValidTickets := validateUserInput(firstUser, lastUser, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstUser+" "+lastUser)

			fmt.Printf("User Name is %v %v and his email address is %v \n", firstUser, lastUser, email)
			fmt.Printf("Remaining Tickets are %v \n", remainingTickets)

			// calling the function to print first name
			firstNames := printFirstName(bookings)
			fmt.Printf("These first name of  bookings :  %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our tickets are booked.Come again next time bye!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Your first and last name is too short , enter correct name")
			}
			if !isValidEmail {
				fmt.Println("Enter correct email! Your email does not contain @ sign")
			}
			if !isValidTickets {
				fmt.Println("Enter valid Number of tickets!. Thank you")
			}
		}

	}

}

// writing function in go language

func greetuser(confName string, remainingTickets uint) {
	fmt.Println("Welcom to my Go tutorial!")
	fmt.Printf("My name is %v \n", confName)
	fmt.Println("Lets enjoy this journey together")
	fmt.Println("Lets make Booking App")
	fmt.Printf("Remainig Tickets are %v \n", remainingTickets)
}

func printFirstName(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstUser string
	var lastUser string
	var email string
	var userTickets uint
	// Taking Input from the user
	fmt.Print("Enter your First Name: \n")
	fmt.Scan(&firstUser)
	fmt.Print("Enter your Last Name: \n")
	fmt.Scan(&lastUser)
	fmt.Print("Enter your email address: \n")
	fmt.Scan(&email)
	fmt.Print("Enter the number of tickets you want to buy : \n")
	fmt.Scan(&userTickets)
	return firstUser, lastUser, email, userTickets
}
