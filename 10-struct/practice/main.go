package main

import "fmt"

type HinhChuNhat struct {
	chieuDai  float32 `desc:"Chieu dai hinh chu nhat"`
	chieuRong float32 `desc:"Chieu rong hinh chu nhat"`
}

func (hcn HinhChuNhat) getChuVi() float32 {
	return (hcn.chieuRong + hcn.chieuDai) * 2

}

func (hcn HinhChuNhat) getDienTich() float32 {
	return hcn.chieuRong * hcn.chieuDai
}

func main() {
	hcn := HinhChuNhat{
		chieuDai:  5,
		chieuRong: 7,
	}
	fmt.Println(hcn.getChuVi())
	fmt.Println(hcn.getDienTich())
}
