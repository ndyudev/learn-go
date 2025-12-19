package main

import "fmt"

// ---------------------------------------------------------
// BÀI TẬP LỚN: MÁY IN HÓA ĐƠN NHÀ HÀNG (RECEIPT PRINTER)
//
// Mục tiêu:
// 1. Khai báo biến chuẩn (int, float64, string).
// 2. Sử dụng toán tử (+, *, /) để tính tiền.
// 3. Dùng fmt.Printf với các định dạng căn lề (%-20s, %10d...) để in bảng đẹp.
//
// DỮ LIỆU ĐẦU VÀO (MENU):
// - Món 1: "Cơm Tấm",   Giá: 35000,  SL: 2
// - Món 2: "Sườn Bì",   Giá: 15000,  SL: 2
// - Món 3: "Trà Đá",    Giá: 5000,   SL: 4
// ---------------------------------------------------------

func main() {
	// ========================================================
	// BƯỚC 1: KHAI BÁO BIẾN (VARIABLES)
	// ========================================================
	// Gợi ý: Tên món dùng string, Giá và SL dùng int
	// TODO: Khai báo tên, giá, số lượng cho 3 món ăn ở đây

	var comTam string = "Cơm Tấm"
	var suonBi string = "Sườn bì"
	var traDa string = "Trà Đá"

	var quantityComTam int = 2
	var quantitySuonBi int = 2
	var quantityTraDa int = 4

	var priceComTam = 35000
	var priceSuonBi = 15000
	var priceTraDa = 5000
	// ========================================================
	// BƯỚC 2: TÍNH TOÁN (LOGIC)
	// ========================================================
	// TODO: Tính thành tiền của từng món (Giá * SL)
	var totalComTam = priceComTam * quantityComTam
	var totalSuonBi = priceSuonBi * quantitySuonBi
	var totalTraDa = priceTraDa * quantityTraDa
	// TODO: Tính tổng tạm tính (Subtotal) = Tổng tiền 3 món cộng lại
	var totalSubBill = totalComTam + totalSuonBi + totalTraDa
	// TODO: Tính thuế VAT (10% của Subtotal)
	// Lưu ý: Subtotal đang là int, muốn nhân 0.1 (float) thì phải ép kiểu: float64(subtotal) * 0.1
	var tax = float64(totalSubBill) * 0.1
	var totalBill = tax + float64(totalSubBill)

	// TODO: Tính TỔNG CỘNG PHẢI TRẢ (Total) = Subtotal + VAT

	// ========================================================
	// BƯỚC 3: IN HÓA ĐƠN (FORMATTING)
	// ========================================================
	// MẪU KẾT QUẢ MONG MUỐN (Cố gắng in giống hệt thế này):
	/*
		=====================================
		          HOA DON THANH TOAN
		=====================================
		Mon an              SL        Don Gia
		-------------------------------------
		Com Tam              2          35000
		Suon Bi              2          15000
		Tra Da               4           5000
		-------------------------------------
		Tam tinh:                      120000
		Thue VAT (10%):                 12000
		-------------------------------------
		TONG CONG:                     132000
		=====================================
	*/

	// TODO: Bắt đầu viết lệnh in ở dưới đây
	fmt.Println("=====================================")
	fmt.Println("          HOA DON THANH TOAN         ")
	fmt.Println("=====================================")
	fmt.Println("Mon an               SL    Thanh Tien")
	fmt.Println("-------------------------------------")
	fmt.Printf("%-17s %5d %13d \n", comTam, quantityComTam, totalComTam)
	fmt.Printf("%-17s %5d %13d \n", suonBi, quantitySuonBi, totalSuonBi)
	fmt.Printf("%-17s %5d %13d \n", traDa, quantityTraDa, totalTraDa)
	fmt.Println("-------------------------------------")
	fmt.Printf("Tam tinh %28d\n", totalSubBill)
	fmt.Printf("Thue VAT (10%%) %22.0f\n", tax)
	fmt.Println("-------------------------------------")
	fmt.Printf("Tong Cong %27.0f", totalBill)

}
