package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age                 int
	gender              bool
}

func showInfomation(e Employee) {
	fmt.Printf("First Name: %s \nLast Name: %s\n Age %d\n", e.firstName, e.lastName, e.age)
}

func main() {

	employee1 := Employee{
		firstName: "Nic",
		lastName:  "Duy",
		gender:    true,
		age:       19,
	}
	fmt.Println(employee1)
	showInfomation(employee1)
}
