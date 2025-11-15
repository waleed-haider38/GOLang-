package user

import "fmt"

type Person struct {
	name   string
	age    uint
	salary uint
}

func Hello() {
	fmt.Println("Hi , I am here to provide you some services")
}

func MyStruct() {
	var pers1 Person
	pers1.name = "Waleed Haider"
	pers1.age = 23
	pers1.salary = 25000

	fmt.Println(pers1.name)
	fmt.Println(pers1.age)
	fmt.Println(pers1.salary)
}
