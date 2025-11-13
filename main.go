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
	userName := "Waleed Haider"
	userAge := 23
	bookings := []string{}

	// different ways to declare the variables

	// var myUser string = "Go Conference"
	// const myage int = 18

	// fmt.Println(userAge)
	// fmt.Println(&userAge)
	// fmt.Printf("userName type is %T Age type is %T \n", userName, userAge)
	fmt.Println("Welcom to my Go tutorial")
	fmt.Printf("My name is %v \n", userName)
	fmt.Printf("I am %v years old\n", userAge)
	fmt.Println("Lets enjoy this journey together")
	fmt.Println("Lets make Booking App")

	var remainingTickets uint = 50
	var userTickets uint

	fmt.Printf("Remainig Tickets are %v \n", remainingTickets)

	for {
		var firstUser string
		var lastUser string
		var email string

		// Taking Input from the user

		fmt.Print("Enter your First Name: \n")
		fmt.Scan(&firstUser)

		fmt.Print("Enter your Last Name: \n")
		fmt.Scan(&lastUser)

		fmt.Print("Enter your email address: \n")
		fmt.Scan(&email)

		fmt.Print("Enter the number of tickets you want to buy : \n")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets and you are trying to book %v tickets", remainingTickets, userTickets)

		}

		remainingTickets = remainingTickets - userTickets

		if remainingTickets == 0 {
			fmt.Println("Our tickets are booked.Come again next tiem bye!")
			break
		} else {
			fmt.Println("Not enough tickets available!")
		}

		bookings = append(bookings, firstUser+" "+lastUser)

		fmt.Printf("User Name is %v %v and his email address is %v \n", firstUser, lastUser, email)
		fmt.Printf("Remaining Tickets are %v \n", remainingTickets)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These first name of  bookings :  %v \n", firstNames)
	}

}
