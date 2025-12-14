package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- For loop --")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}
