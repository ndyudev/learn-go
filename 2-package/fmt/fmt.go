package main

import "fmt"

func main() {
	var fullName = "ndyudev"
	age := 18
	fmt.Print("Hello world! \n") // print này cần \n mới xuống dòng được
	fmt.Println("Hello world!")
	fmt.Println("Goodbye world!")

	fmt.Printf("Hello %s and age : %d \n", fullName, age)
	var firstName, lastName string
	var address string
	fmt.Println("Nhập họ và ten của bạn:")
	//fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Scan(&firstName, &lastName)
	//fmt.Scanln(&firstName, &lastName, &address)
	fmt.Println("Nhập địa chỉ của bạn:")
	fmt.Scan(&address)
	fmt.Printf("FirstName: %s LastName: %s \n Address: %s \n", firstName, lastName, address)

	var school = "UIT"
	edu := fmt.Sprintf("University of Information Technology! | %s ", school)
	fmt.Println(edu)

	var city string
	var street string
	var District string
	fmt.Println("Nhập đường quận và thành phố của bạn:")
	fmt.Scan(&street, &District, &city)
	addressDetail := fmt.Sprintf("Street: %v \n District: %s \n, City: %s", street, District, city)
	fmt.Println(addressDetail)
}
