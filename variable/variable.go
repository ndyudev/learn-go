package main

import (
	"fmt"
)

var address = "Ho Chi Minh City"

var (
	monhoc   string
	thongtin string
) // khai báo nhiều biến bên ngoài

//course := "Go Web Programming Language" // biến ngoài hàm thì nên khai báo var vì := vô hàm kh xài được

const FULL_NAME = "CHAU Nhat Duy"

func main() {
	// integer literal
	fmt.Println(
		-200, -100, 0, 50, 100, 100,
	)

	// float literal
	fmt.Println(
		-50.5, -20.5, -0., 1., 100.56, // ...
	)

	// bool constants
	fmt.Println(
		true, false,
	)

	// string literal - utf-8
	fmt.Println(
		"", // chuỗi rỗng sẽ thành dấu cách
		"hi",

		// unicode
		"nasılsın?",         // "how are you" in turkish
		"hi there ndyudev!", // "hi there star!"
	)
	var fullName string = "Chau Nhat Duy" // bỏ kiểu dữ liệu thì lấy từ value - bỏ string thì vẫn nhận biết là string
	fullname := fullName
	fmt.Println(fullName)
	fmt.Println(fullname)
	var age int // khai báo biến trước - không gán là 0
	age = 20    // gán sau
	fmt.Println(age)

	phone := "0125421" // := thì không gán kiểu dữ liệu và bắt buộc có value
	fmt.Println(phone)

	fmt.Println(address)
	var toan, ly, hoa int // var thì khai báo được nhiều
	toan = 9
	ly = 10
	hoa = 8
	fmt.Println(toan, ly, hoa)

	van, su, dia := 8, 6, 9 // phải khai báo trên 1 dòng
	fmt.Println(van, su, dia)

	monhoc = "Lập trình go lang"

	fmt.Println(monhoc)

	thongtin = "ndyudev"

	fmt.Println(thongtin)
	var hoTen string
	fmt.Print("Vui lòng nhập họ tên:")
	//fmt.Scanln(&hoTen) // Nhập họ tên từ scanln - space hoặc enter thì nó xong việc nhập luôn
	//scanner := bufio.NewScanner(os.Stdin) // khởi tạo scanner bufio và os lirary
	//if scanner.Scan() {
	//	hoTen = scanner.Text()
	//}
	fmt.Println("Ho ten:", hoTen)
	// Hằng số - lưu trữ 1 giá kh được ghi đè cùng 1 phạm vi - nên VIẾT_HOA_TOÀN_BỘ - khai báo trong và ngoài hàm
	//const FULL_NAME = "Chau Nhat Duy"
	fmt.Println(FULL_NAME)
}
