package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age                 int
	gender              bool
}

func main() {

	employee1 := Employee{
		firstName: "Nic",
		lastName:  "Duy",
		gender:    true,
		age:       19,
	}

	fmt.Println(employee1)
}
