package main

import "fmt" //fmt is package where we can use some built in functions like Print , Println etc there are multiple package which we see next in our tutorial.
// print is used to print some statement but println is also used to print statement but it starts a new line. Now we use printf function which is used to identify the varialbes in effiecient way.
// %v is used to show the value of variable but %T is used to show the type of variable.
// := is only used with var and we do not need to explicity identify the type of variable so it save our time.

func main() {
	userName := "Waleed Haider"
	userAge := 23

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

	// Taking Input from the user

	var firstUser string
	var lastUser string
	var email string

	fmt.Print("Enter your First Name: \n")
	fmt.Scan(&firstUser)

	fmt.Print("Enter your Last Name: \n")
	fmt.Scan(&lastUser)

	fmt.Print("Enter your email address: \n")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want to buy : \n")
	fmt.Scan(&userTickets)

	if userTickets <= remainingTickets {
		remainingTickets = remainingTickets - userTickets
	} else {
		fmt.Println("Not enough tickets available!")
	}

	fmt.Printf("User Name is %v %v and his email address is %v \n", firstUser, lastUser, email)
	fmt.Printf("Remaining Tickets are %v", remainingTickets)

}
