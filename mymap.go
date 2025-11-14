package main

import "fmt"

var firstName string
var lastName string
var userData = make(map[string]string)

func myapp() {

	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter Your Last Name: ")
	fmt.Scan(&lastName)

	userData["firstname"] = firstName
	userData["lastname"] = lastName
	fmt.Println(userData)
}
