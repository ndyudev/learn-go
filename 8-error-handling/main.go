package main

import "fmt"

// * Viết hàm chia 2 số, bắt lỗi chia cho 0
// * Viết hàm kiểm tra tuổi (tuổi < 18 → trả về lỗi)
// * In lỗi ra console
func divisionTwo(numb1, numb2 float64) (sum float64, err error) {
	if numb2 == 0 {
		err = fmt.Errorf("lỗi nguy hiểm: không thể chia cho số 0")
		return
	}
	sum = numb1 / numb2

	return
}

func checkAge(age int) (result bool, err error) {
	if age < 18 {
		err = fmt.Errorf("Tuổi phải lớn đủ 18")
		result = false
		return
	}
	result = true
	return
}

func main() {
	fmt.Println("-- Errors Handling Golang --")

	ketQua, loi := divisionTwo(6, 0)

	if loi != nil {
		fmt.Println("Gặp lỗi rồi:", loi)
	} else {
		fmt.Println("Kết quả phép chia:", ketQua)
	}

	ketQua2, loi2 := divisionTwo(10, 2)
	if loi2 == nil {
		fmt.Println("Kết quả phép chia 2:", ketQua2)
	}

	isAdult, errorAge := checkAge(17)

	if isAdult == true && errorAge == nil {
		fmt.Println("Đủ tuổi!")
	} else {
		fmt.Println("Không đủ tuổi!", errorAge)
	}
}
