package main

import (
	"fmt"
)

func main() {
	// Bài 1: Lặp từ 1 -> 100, bỏ qua 6,48, 75, 89 and mỗi số cách nhau bằng dấu phải ỏ cuối thì thêm dấu chấm
	for i := 1; i <= 100; i++ {
		switch i {
		case 6, 48, 75, 89:
			continue
		}
		if i == 100 {

			fmt.Print(100, ".")
			continue
		}
		fmt.Print(i, ",")
	}
	// Bài 2: loop 1->100 , hiển thị số lẻ, cứ mỗi 3 số thì xuống dòng và không có dấu phẩy cuối dòng,
	//mỗi số cách nau bằng dấu phẩy bỏ ở vị trí cuối cùng
	var count = 0
	fmt.Println()
	for i2 := 1; i2 <= 100; i2++ {
		if i2%2 != 0 {
			fmt.Print(i2)
			count++
			if i2 == 99 {
				fmt.Print(".")
				break
			}

			if count == 3 {
				fmt.Println()
				count = 0
			} else {

				fmt.Print(",")
			}
		}
	}
	// Bài 3: Hiển thị bảng cửu chương
	// cho phép người dùng nhập số bắt đầu và kết thúc kiểm tra điều kiện input nếu người dùng nhập sai.
	// Mỗi bảng cửu chương có tiêu đề hiển thị.
	var cuuchuong, a, b int
	fmt.Println("Nhập số bảng cửu chương bạn muốn in: ")
	fmt.Scanln(&cuuchuong)
	fmt.Println("Nhập số bắt đầu:")
	fmt.Scanln(&a)
	fmt.Println("Nhập số kết thúc:")
	fmt.Scanln(&b)
	fmt.Printf("Bảng cửu chương %d \n", cuuchuong)
	if a > b {
		fmt.Println("Lỗi: Số bắt đầu không được lớn hơn số kết thúc!")
		return
	}
	for i3 := a; i3 <= b; i3++ {
		fmt.Printf("%d * %d = %d\n", cuuchuong, i3, cuuchuong*i3)
	}
}
