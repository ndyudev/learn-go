package main

import "fmt"

// Bài tập 1: Xây dựng 1 hàm đệ quy tính tổng dãy N từ ngươì dùng nhập
// VD: Nhập 10 thì nhập tính tổng từ 1 tới 10
func sumDigi(numb1 int) int {
	if numb1 == 0 {
		return numb1
	}
	return numb1 + sumDigi(numb1-1)
}

// Bài tập 2: Xây dựng 1 hàm để hiển thị dãy số Fibonacci
// - Hiển thị dãy số Fibonacc và tính tổng dãy số đó.
func fibonacci(numb int) {
	term1, term2 := 0, 1
	var total int
	if numb == 0 {
		return
	}
	if numb == 1 {
		fmt.Printf("Day so fibonacci: %d\n", term1)
	} else if numb == 2 {
		fmt.Printf("Day so fibonacci: %d %d\n", term1, term2)
	} else {
		total = term1 + term2
		fmt.Printf("Day so fibonacci: %d %d", term1, term2)
		for i := 3; i <= numb; i++ {
			term3 := term1 + term2
			fmt.Printf(" %d", term3)
			total += term3
			term1, term2 = term2, term3
		}
		fmt.Printf("\nTong day so fibonacci: %d\n", total)
	}
	return
}

func main() {
	var choice int
	for {
		fmt.Println()
		fmt.Println("-=-=-=-==-= MENU CHỨC NĂNG -=-=-=-==-=")
		fmt.Println("[1] Tính tổng dãy số bằng đệ quy")
		fmt.Println("[2] Hiển thị và tính tổng dãy số Fibonacci bằng đệ quy ")
		fmt.Println("[0] Thoát chương trình")
		fmt.Println("-- Vui lòng nhập chức năng số của chức chức năng:")
		fmt.Println("-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var numb int
			fmt.Println("Nhập số muốn tổng")
			_, err := fmt.Scan(&numb)
			if err != nil || numb <= 0 {
				fmt.Printf("Vui long nhap so nguyen duong!\n")
				continue
			}
			var result = sumDigi(numb)
			fmt.Println(result)
		case 2:
			var fibo int
			fmt.Println("Nhập độ dài dãy số Fibonacci:")
			_, err := fmt.Scan(&fibo)
			if fibo <= 0 || err != nil {
				fmt.Printf("Đồ dài dãy số Fibonacci phải lớn hơn 0\n")
				continue
			}
			fibonacci(fibo)
		case 0:
			fmt.Println("Đang thoát chương trình... Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ!")
		}
	}
}
