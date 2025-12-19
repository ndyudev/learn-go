# Cheat Sheet: Toán tử & Phép gán trong Go

Tổng hợp các toán tử cơ bản cần nhớ khi lập trình Go.

## 1. Toán tử Số học (Arithmetic Operators)

Dùng để tính toán số liệu.

| Toán tử | Mô tả | Ví dụ (`a := 10`, `b := 3`) | Kết quả |
| :--- | :--- | :--- | :--- |
| **`+`** | Cộng | `a + b` | `13` |
| **`-`** | Trừ | `a - b` | `7` |
| **`*`** | Nhân | `a * b` | `30` |
| **`/`** | Chia lấy nguyên | `a / b` | `3` (Mất phần thập phân) |
| **`%`** | Chia lấy dư (Modulus) | `a % b` | `1` |

> **⚠️ Lưu ý quan trọng:**
> Trong Go, nếu chia 2 số nguyên (`int`), kết quả sẽ luôn là `int`.
> * `10 / 3` -> `3` (Không phải 3.33)
> * Muốn ra số thực, phải ép kiểu: `float64(10) / 3` -> `3.333...`

---

## 2. Toán tử Tăng/Giảm (Increment/Decrement)

Đây là điểm **KHÁC BIỆT LỚN** so với Java/C++.

| Toán tử | Mô tả | Ví dụ |
| :--- | :--- | :--- |
| **`++`** | Tăng thêm 1 đơn vị | `i++` (Tương đương `i = i + 1`) |
| **`--`** | Giảm đi 1 đơn vị | `i--` (Tương đương `i = i - 1`) |

### ❌ Những điều CẤM KỴ trong Go:
1.  **Không có tiền tố (Prefix):** Không được viết `++i` hay `--i`. Chỉ được viết `i++`.
2.  **Không phải là biểu thức:** Không được gán hoặc in trực tiếp.
    ```go
    // Java: int x = i++; (OK)
    // Go:   x := i++     (LỖI Syntax ngay lập tức)
    
    // Java: System.out.println(i++); (OK)
    // Go:   fmt.Println(i++)         (LỖI)
    ```
    *Cách đúng:* Tăng `i++` ở dòng riêng biệt, rồi mới dùng `i`.

---

## 3. Toán tử So sánh (Comparison Operators)

Kết quả trả về luôn là `bool` (`true` hoặc `false`).

| Toán tử | Mô tả | Ví dụ (`10` vs `20`) |
| :--- | :--- | :--- |
| **`==`** | Bằng nhau | `10 == 20` -> `false` |
| **`!=`** | Khác nhau | `10 != 20` -> `true` |
| **`>`** | Lớn hơn | `10 > 20` -> `false` |
| **`<`** | Nhỏ hơn | `10 < 20` -> `true` |
| **`>=`** | Lớn hơn hoặc bằng | `10 >= 10` -> `true` |
| **`<=`** | Nhỏ hơn hoặc bằng | `5 <= 10` -> `true` |

---

## 4. Toán tử Logic (Logical Operators)

Dùng để kết hợp nhiều điều kiện (`if`, `for`).

| Toán tử | Tên | Mô tả | Ví dụ |
| :--- | :--- | :--- | :--- |
| **`&&`** | AND (Và) | Cả 2 đều đúng mới là đúng | `true && false` -> `false` |
| **`\|\|`** | OR (Hoặc) | Chỉ cần 1 cái đúng là đúng | `true \|\| false` -> `true` |
| **`!`** | NOT (Phủ định) | Đổi đúng thành sai, sai thành đúng | `!true` -> `false` |

---

## 5. Toán tử Gán (Assignment Operators)

Dùng để lưu giá trị vào biến.

| Toán tử | Ví dụ | Tương đương với | Giải thích |
| :--- | :--- | :--- | :--- |
| **`=`** | `x = 10` | | Gán giá trị thông thường |
| **`:=`** | `x := 10` | `var x = 10` | Khai báo + Gán (Chỉ trong hàm) |
| **`+=`** | `x += 5` | `x = x + 5` | Cộng thêm |
| **`-=`** | `x -= 5` | `x = x - 5` | Trừ bớt |
| **`*=`** | `x *= 5` | `x = x * 5` | Nhân dồn |
| **`/=`** | `x /= 5` | `x = x / 5` | Chia dồn |

---

## 6. Bitwise Operators (Toán tử Bit) - Tham khảo
(Dành cho xử lý nhị phân, ít dùng trong web cơ bản)

* `&` (AND), `|` (OR), `^` (XOR), `<<` (Dịch trái), `>>` (Dịch phải).