
Câu 1
Which one below is an integer literal? (Cái nào dưới đây là một literal số nguyên?)

✅ -42 (CORRECT - ĐÚNG)

English: This is an integer literal.

Tiếng Việt: Đây là một literal số nguyên (số nguyên âm).

"Hello"

English: This is a string literal.

Tiếng Việt: Đây là một literal chuỗi (nằm trong ngoặc kép).

false

English: This is a bool constant.

Tiếng Việt: Đây là một hằng số boolean (logic đúng/sai).

3.14

English: This is a float literal.

Tiếng Việt: Đây là một literal số thực (có dấu chấm động).

Câu 2
Which one below is a float literal? (Cái nào dưới đây là một literal số thực?)

-42 (Số nguyên / Integer)

"Hello" (Chuỗi / String)

false (Boolean)

✅ 3.14 (CORRECT - ĐÚNG)

Giải thích: Số thực trong Go luôn dùng dấu chấm . để ngăn cách phần thập phân.

Câu 3 (Câu này dễ nhầm lẫn nhất)
Which one below is a float literal? (Cái nào dưới đây là một literal số thực hợp lệ?)

6,28

Sai: Trong lập trình, dấu phẩy , không dùng để viết số thập phân (khác với toán học ở Việt Nam).

,28

Sai: Bắt đầu bằng dấu phẩy là lỗi cú pháp (Illegal).

✅ .28 (CORRECT - ĐÚNG)

Giải thích: Đây là cách viết tắt của 0.28. Go hiểu .28 chính là một số thực float.

628

Sai: Đây là số nguyên (Integer), vì không có dấu chấm.

Câu 4
Which one below is a string literal? (Cái nào dưới đây là một literal chuỗi?)

-42

✅ "Hello" (CORRECT - ĐÚNG)

Giải thích: Bất cứ thứ gì nằm trong cặp dấu ngoặc kép "" đều là String Literal.

false

3.14

Câu 5
Which one below is a bool constant? (Cái nào dưới đây là một hằng số boolean?)

-42

"Hello"

✅ false (CORRECT - ĐÚNG)

Giải thích: Kiểu bool trong Go chỉ có 2 giá trị duy nhất: true (đúng) hoặc false (sai). Chúng được gọi là hằng số (constant) vì giá trị này được định nghĩa sẵn, không thay đổi được.

3.14

# Ghi chú bài học: Biến trong Go (Variables)

Tài liệu tóm tắt kiến thức về cách khai báo và sử dụng biến trong Golang.

## 1. Bảng so sánh: `var` vs `:=`

| Đặc điểm | `var` (Truyền thống) | `:=` (Short Declaration) |
| :--- | :--- | :--- |
| **Cú pháp** | `var x int = 10` hoặc `var x int` | `x := 10` |
| **Phạm vi (Scope)** | Dùng được **MỌI NƠI** (Trong & Ngoài hàm) | Chỉ dùng **TRONG HÀM** |
| **Kiểu dữ liệu** | Có thể chỉ định rõ (`int`, `int64`, `float32`) | Máy tự đoán dựa vào giá trị (Type Inference) |
| **Giá trị khởi tạo** | Không bắt buộc (tự nhận Zero Value) | **Bắt buộc** phải có giá trị ngay |

---

## 2. Khi nào dùng `var`?
Cách khai báo `var` giống với Java/C++. Nên dùng trong 3 trường hợp:

1.  **Khai báo biến toàn cục (Global Variable):**
    Biến dùng chung cho cả package, nằm ngoài hàm `main` hoặc các hàm khác.
    ```go
    var Version = "1.0.0" // Bắt buộc dùng var
    // Version := "1.0.0" // -> LỖI Syntax
    ```

2.  **Khai báo trước, gán giá trị sau:**
    ```go
    var email string
    // ... xử lý logic ...
    email = "admin@gmail.com"
    ```

3.  **Muốn ép kiểu dữ liệu cụ thể:**
    Go mặc định số thực là `float64`. Nếu muốn dùng `float32` để nhẹ máy hơn:
    ```go
    var diemTB float32 = 8.5
    ```

---

## 3. Khi nào dùng `:=` (Short Declaration)?
Đây là cách code chuẩn (Idiomatic Go) bên trong các hàm xử lý.

* **Cơ chế:** Trình biên dịch tự nhìn giá trị bên phải để gán kiểu cho biến bên trái.
* **Ví dụ:**
    ```go
    func main() {
        name := "Gemini" // Go tự hiểu là string
        age := 20        // Go tự hiểu là int
    }
    ```
* **Quy tắc "Biến mới":** Bên trái dấu `:=` phải có ít nhất một biến chưa từng được khai báo.

---

## 4. Các khái niệm quan trọng khác

### A. Zero Values (Giá trị mặc định)
Trong Go, biến khai báo mà chưa gán giá trị sẽ **không bị null** hay rác, mà nhận giá trị mặc định:
* `int`, `float`: `0`
* `string`: `""` (Chuỗi rỗng)
* `bool`: `false`
* `pointer`, `interface`: `nil`

### B. Quy tắc biến thừa (Unused Variables)
Go rất nghiêm khắc. Nếu khai báo biến cục bộ mà **không sử dụng**, chương trình sẽ báo lỗi và **không cho chạy**.
* *Cách xử lý:* Xóa biến đó hoặc gán vào `_` (Blank Identifier) nếu bắt buộc phải khai báo.

### C. Literal vs Illegal
* **Literal:** Giá trị thô được viết trực tiếp vào code (VD: `10`, `"Hello"`, `true`).
* **Illegal:** Code vi phạm cú pháp (VD: Dùng `:=` ngoài hàm, dùng dấu nháy cong `”` thay vì nháy thẳng `"`).