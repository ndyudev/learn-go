package main

import "fmt"

func calculator() {
	var a, b float64
	var pheptinh string
	var result float64

	fmt.Println("-- Chào mứng bạn đến với chương trình tính toán!")
	fmt.Println("Nhập giá trị của a:")
	fmt.Scanln(&a)
	fmt.Println("Nhập giá trị của b:")
	fmt.Scanln(&b)
	fmt.Printf("Vui lòng nhập phép toán bạn muuốn thực hiện giữa 2 số: %.2f và %.2f: \n", a, b)
	fmt.Scanln(&pheptinh)
	switch pheptinh {
	case "+":
		result = a + b
	case "-":
		result = a - b
		//fallthrough ép chạy xuống case ở dưới ! giống continue ở java hoặc js
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("B bẳng 0 không thể chia!")
			return
		}
		result = a / b
	default:
		fmt.Println("Phép tính không hợp lệ 4 phép tính cộng  : +, trừ: -, nhân: *, chia: /")
		return
	}

	fmt.Println(result)
}
