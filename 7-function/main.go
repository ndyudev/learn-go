package main

import "fmt"

func phepToan(a, b int) (tong int) {
	tong = a + b
	return tong
}

func main() {
	fmt.Println("-- Functions Golang --")
	fmt.Print(phepToan(1, 2))
}
