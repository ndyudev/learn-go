package main

import "fmt"

func main() {
	var point = 4.1

	if point < 0 || point > 10 {
		fmt.Println("Điểm không hợp lệ")
		return
	}
	if point >= 9 {
		fmt.Println("Xuất sắc")
	} else if point >= 8 {
		fmt.Println("Học sinh giỏi")
	} else if point >= 7 {
		fmt.Println("Học sinh khá")
	} else if point >= 6 {
		fmt.Println("Học sinh trung bình")
	} else if point >= 5 {
		fmt.Println("Học sinh yếu")
	} else {
		fmt.Println("Học sinh kém")
	}
}
