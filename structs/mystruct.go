package main

import (
	"fmt"
)

func main() {
	MyStruct()

	var greet []string

	for i := 1; i <= 4; i++ {
		var userName, secName string
		fmt.Print("Hi, Enter your first Name: ")
		fmt.Scan(&userName)
		fmt.Print("Enter your second Name: ")
		fmt.Scan(&secName)

		fullName := userName + " " + secName
		greet = append(greet, fullName)
	}

	fmt.Println("The new members are:")
	for _, name := range greet {
		fmt.Println(name)
	}
}

func MyStruct() {
	fmt.Println("Here we will discuss about struct")
}
