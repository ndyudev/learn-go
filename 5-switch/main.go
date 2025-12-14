package main

import "fmt"

func main() {
	fmt.Println("switch case")
	var choice = 5
	switch choice {
	case 1, 5:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Không hợp lệ")
	}
}
