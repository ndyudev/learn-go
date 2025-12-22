package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age                 int
	gender              bool
}

func (e Employee) showInfomation() {
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
	employee1.showInfomation()
}
