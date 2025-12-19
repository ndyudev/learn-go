package main

import "fmt"

// 🧩 Bài 1: Basic Pointer
// Khai báo biến number
// Tạo pointer trỏ tới number
// Thay đổi giá trị của number thông qua pointer
func changeValue(number int) {
	age := number
	ptnAge := &age

	*ptnAge = 19
	fmt.Println("Address number:", &number)
	fmt.Println("Address age:", &age)
	fmt.Println("Value age:", age)
	fmt.Println("Value number:", number)
}

// 🧩 Bài 2: Swap Function (Hoán đổi)
func swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println("Value A:", *a)
	fmt.Println("Value B:", *b)
}

func main() {
	fmt.Println("-- Practice Pointers --")
	changeValue(18)
	a := 19
	b := 18
	swap(&a, &b)
}
