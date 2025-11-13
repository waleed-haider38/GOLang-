package main

import "fmt" //fmt is package where we can use some built in functions like Print , Println etc there are multiple package which we see next in our tutorial.
// print is used to print some statement but println is also used to print statement but it starts a new line. Now we use printf function which is used to identify the varialbes in effiecient way.
// %v is used to show the value of variable but %T is used to show the type of variable.
// := is only used with var and we do not need to explicity identify the type of variable so it save our time.

func main() {
	userName := "Waleed Haider"
	userAge := 23
	// fmt.Printf("userName type is %T Age type is %T \n", userName, userAge)
	fmt.Println("Welcom to my Go tutorial")
	fmt.Printf("My name is %v \n", userName)
	fmt.Printf("I am %v years old\n", userAge)
	fmt.Println("Lets enjoy this journey together")

	// fmt.Println(userAge)
	// fmt.Println(&userAge)

	// Taking Input from the user

	var user string

	fmt.Print("Enter your First Name: \n")
	fmt.Scan(&user)

	fmt.Printf("User Name is %v", user)

}
