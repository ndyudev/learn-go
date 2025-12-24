package main

import "fmt"

type Student struct {
	firstName string `json:"first_name,omitempty"`
	lastName  string `json:"last_name,omitempty"`
	age       int    `json:"age,omitempty"`
	gender    bool   `json:"gender,omitempty"`
}

func (s *Student) showInformation() {
	fmt.Println(s.firstName, s.lastName, s.age, s.gender)
}

func (s *Student) deleteStudent() {
	s.firstName = ""
	s.lastName = ""
	s.age = 0
	s.gender = false
}

func main() {
	fmt.Println("--- Struct Tag ---")

	student1 := Student{
		firstName: "Duy",
		lastName:  "Chau Nhat",
		age:       19,
		gender:    true,
	}

	student1.showInformation()
	student1.deleteStudent()
	student1.showInformation()
}
