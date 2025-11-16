package main

import (
	"fmt"
	"time"
)

type School struct {
	school_name  string
	teacher_name string
	student_name string
	roll_no      uint
}

func main() {
	go MyStruct()

	time.Sleep(1 * time.Second)

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
	var std1 School
	std1.school_name = "Scholar Public School"
	std1.teacher_name = "Sir Sohail"
	std1.student_name = "Muhammad Waleed Haider"
	std1.roll_no = 11

	fmt.Println(std1.school_name)
	fmt.Println(std1.teacher_name)
	fmt.Println(std1.student_name)
	fmt.Println(std1.roll_no)
}
