package main

import "fmt"

func main() {
	fmt.Println("-- Formating Verbs --")
	ten := "ndyudev"
	age := 20
	height := 1.75
	isMale := false

	fmt.Printf("Name: %s\n", ten)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.1fm \n", height) // nếu làm tròn thì ít hơn 1 chữ số 1.75 = %.1f => 1m8
	fmt.Printf("Male: %t\n", isMale)
	fmt.Printf("Name: %v \nAge: %v \nHeight: %.3v\nMale: %v", ten, age, height, isMale) // %v giúp lấy số dựa vào giá trị trị vào float thì nhiều hơn 1 số
}
