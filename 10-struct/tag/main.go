package main

import "fmt"

type Student struct {
	firstName string `json:"first_name,omitempty"`
	lastName  string `json:"last_name,omitempty"`
	age       int    `json:"age,omitempty"`
	gender    bool   `json:"gender,omitempty"`
}

func main() {
	fmt.Println("--- Struct Tag ---")

}
