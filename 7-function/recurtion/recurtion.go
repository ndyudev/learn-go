package main

import "fmt"

func phepToan(numb1, numb2 int) (int, int, int, float32) {
	if numb2 == 0 {
		numb2 = 1
	}

	tong := numb1 + numb2
	tru := numb1 - numb2
	nhan := numb1 * numb2
	chia := float32(numb1) / float32(numb2)
	return tong, tru, nhan, chia
}

func giaiThua(numb int) int {
	if numb == 0 {
		return 1
	}
	return numb * giaiThua(numb-1)
}

func countDown(number int) {
	fmt.Println(number)
	if number > 0 {
		countDown(number - 1)
	}
}

func main() {
	cong, tru, nhan, _ := phepToan(2012, 2)

	fmt.Println(cong, tru, nhan)
	countDown(10)

	ketQua := giaiThua(5)
	fmt.Println("Giai thừa của 5 là:", ketQua)

}
