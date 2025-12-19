package main

import "fmt"

func main() {
	fmt.Println("--- Golang Pointers ---")
	age := 17

	// 1. Biến thường
	fmt.Printf("Data type age: %T \n", age)
	fmt.Printf("Value age: %d \n", age)
	fmt.Println("Address Age:", &age)

	fmt.Println("--------------------")

	// 2. Biến con trỏ
	ptnAge := &age
	ptnAge2 := &ptnAge

	fmt.Printf("Data type ptnAge: %T \n", ptnAge) // -> *int

	fmt.Printf("Value ptnAge (Address): %p \n", ptnAge)

	fmt.Printf("Value inside ptnAge (Dereference): %d \n", *ptnAge)

	fmt.Println("Address ptnAge:", &ptnAge)

	fmt.Printf("Data type ptnAge2: %T \n", ptnAge2) // -> *int

	fmt.Printf("Value ptnAge2 (Address): %p \n", ptnAge2)

	fmt.Printf("Value inside ptnAge2 (Dereference): %d \n", **ptnAge2)

	fmt.Println("Address ptnAge2:", &ptnAge2)
}
