# Cheat Sheet: Formatting Verbs (fmt)

Tài liệu tra cứu nhanh các ký tự định dạng trong hàm `fmt.Printf()` và `fmt.Sprintf()`.

## 1. Verbs Phổ Biến Nhất (General)

Những ký tự này dùng được cho hầu hết mọi kiểu dữ liệu.

| Verb | Mô tả | Ví dụ Code | Kết quả |
| :--- | :--- | :--- | :--- |
| **`%v`** | **Value**. In giá trị mặc định. | `Printf("%v", 10)` | `10` |
| **`%T`** | **Type**. In ra kiểu dữ liệu của biến. | `Printf("%T", 10)` | `int` |
| **`%%`** | In ra dấu phần trăm thực sự. | `Printf("%d%%", 50)` | `50%` |

---

## 2. Số Nguyên (Integer)

Dùng cho các kiểu `int`, `int8`, `int64`, `uint`...

| Verb | Mô tả | Ví dụ (`n := 255`) | Kết quả |
| :--- | :--- | :--- | :--- |
| **`%d`** | **Decimal**. Hệ thập phân (cơ số 10). | `Printf("%d", n)` | `255` |
| **`%b`** | **Binary**. Hệ nhị phân (cơ số 2). | `Printf("%b", n)` | `11111111` |
| **`%x`** | **Hex**. Hệ 16 (chữ thường). | `Printf("%x", n)` | `ff` |
| **`%X`** | **HEX**. Hệ 16 (chữ HOA). | `Printf("%X", n)` | `FF` |

---

## 3. Số Thực (Float)

Dùng cho `float32`, `float64`.

| Verb | Mô tả | Ví dụ (`f := 12.3456`) | Kết quả |
| :--- | :--- | :--- | :--- |
| **`%f`** | **Float**. Mặc định lấy 6 số lẻ. | `Printf("%f", f)` | `12.345600` |
| **`%.2f`** | **Precision**. Chỉ lấy 2 số lẻ (làm tròn). | `Printf("%.2f", f)` | `12.35` |
| **`%e`** | Ký hiệu khoa học (Scientific). | `Printf("%e", f)` | `1.234560e+01` |

---

## 4. Chuỗi & Ký Tự (String & Rune)

| Verb | Mô tả | Ví dụ | Kết quả |
| :--- | :--- | :--- | :--- |
| **`%s`** | **String**. In chuỗi bình thường. | `Printf("%s", "Hello")` | `Hello` |
| **`%q`** | **Quoted**. In chuỗi kèm dấu ngoặc kép. | `Printf("%q", "Hello")` | `"Hello"` |
| **`%c`** | **Character**. In ký tự từ mã Unicode (Rune). | `Printf("%c", 65)` | `A` |

---

## 5. Boolean & Pointer

| Verb | Mô tả | Ví dụ | Kết quả |
| :--- | :--- | :--- | :--- |
| **`%t`** | **True/False**. | `Printf("%t", true)` | `true` |
| **`%p`** | **Pointer**. In địa chỉ bộ nhớ (Hex, bắt đầu 0x). | `Printf("%p", &x)` | `0xc0000...` |

---

## 6. Căn Chỉnh & Độ Rộng (Width & Padding)

Dùng để in bảng biểu cho thẳng hàng.

* **`%9f`**: Độ rộng tối thiểu là 9 ký tự (căn lề phải).
* **`%.2f`**: Lấy 2 số thập phân.
* **`%9.2f`**: Độ rộng 9 ký tự VÀ lấy 2 số thập phân.
* **`%-9d`**: Dấu trừ `-` nghĩa là căn lề TRÁI.

**Ví dụ thực tế:**
```go
fmt.Printf("|%6d|%6d|\n", 12, 345)
fmt.Printf("|%6d|%6d|\n", 1234, 5)