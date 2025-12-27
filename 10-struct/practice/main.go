package main

import "fmt"

type HinhChuNhat struct {
	chieuDai  float32 `desc:"Chieu dai hinh chu nhat"`
	chieuRong float32 `desc:"Chieu rong hinh chu nhat"`
}

func (hcn *HinhChuNhat) getChuVi() float32 {
	return (hcn.chieuRong + hcn.chieuDai) * 2

}

func (hcn *HinhChuNhat) getDienTich() float32 {
	return hcn.chieuRong * hcn.chieuDai
}

func main() {
	var hinhChuNhat HinhChuNhat

	for {
		fmt.Println("Nhập chiều dài hình chữ nhật:")
		_, err := fmt.Scanf("%f", &hinhChuNhat.chieuDai)

		if err == nil && hinhChuNhat.chieuDai > 0 {
			break
		}
		fmt.Println("Kích thước chiều rộng phải lớn hơn 0 !")

	}

	for {
		fmt.Println("Nhập chiều rộng hình chữ nhật:")
		_, err := fmt.Scanf("%f", &hinhChuNhat.chieuRong)

		if err == nil && hinhChuNhat.chieuRong > 0 {
			break
		}
		fmt.Println("Kích thước chiều dài phải lớn hơn 0 !")
	}
	fmt.Println(hinhChuNhat.getChuVi())
	fmt.Println(hinhChuNhat.getDienTich())
}
