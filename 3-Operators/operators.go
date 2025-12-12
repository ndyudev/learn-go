package main

import "fmt"

func main() {
	fmt.Println("-- Operators --")
	var s1 = 5
	var s2 = 12
	tong := s1 + s2
	hieu := s1 - s2
	tich := s1 * s2
	thuong := float32(s1) / float32(s2) // ép kiểu sang float trước rồi chia mới trả về float.
	du := s2 % s1
	du--
	tong++
	fmt.Println(tong)
	fmt.Println(hieu)
	fmt.Println(tich)
	fmt.Println(thuong)
	fmt.Println(du)
	fmt.Println(s1 >= 6 || s2 > 6)
	isTrue := !(s1 > 6)
	fmt.Println(isTrue)
	s2++
	s2 -= 5
	fmt.Println(s2)
}
