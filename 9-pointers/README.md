# 9. Pointers in Go (Con trỏ)

Trong Go, **mọi thứ mặc định đều là *Pass by Value*** (truyền tham trị – copy giá trị).
Muốn **thay đổi dữ liệu gốc** hoặc **tối ưu hiệu năng**, ta cần dùng **Pointers (Con trỏ)**.

> 👉 Pointer là một trong những khái niệm **cốt lõi nhất** của Go. Hiểu được pointer là coi như lên trình hẳn 😎

---

## 1️⃣ Khái niệm cốt lõi (Concept)

Hãy tưởng tượng **RAM** giống như một khu phố:

* **Biến (Variable)** 👉 một ngôi nhà, bên trong chứa dữ liệu (ví dụ: `10`, `"Hello"`)
* **Con trỏ (Pointer)** 👉 tờ giấy ghi **địa chỉ ngôi nhà** đó (ví dụ: `0xC00004`)

📌 **Kết luận:**

> Pointer là biến dùng để lưu **ĐỊA CHỈ BỘ NHỚ** của một biến khác.

---

## 2️⃣ Hai ký hiệu "thần thánh" cần nhớ

| Ký hiệu | Tên gọi     | Ý nghĩa                     | Ví dụ               |
| :-----: | :---------- | :-------------------------- | :------------------ |
| **`&`** | Address-of  | Lấy **địa chỉ** của biến    | `&a` → `0xc00001`   |
| **`*`** | Dereference | Lấy **giá trị tại địa chỉ** | `*p` → giá trị thật |

👉 Nhớ nhanh:

* `&` = **xem nhà ở đâu** 🏠
* `*` = **vào nhà lấy đồ** 📦

---

## 3️⃣ Khai báo và sử dụng Pointer

```go
package main

import "fmt"

func main() {
    var a int = 100

    // 1. Tạo con trỏ p trỏ tới a
    var p *int = &a

    // 2. In thông tin
    fmt.Println("Giá trị của a:", a)
    fmt.Println("Địa chỉ của a:", &a)
    fmt.Println("Giá trị của p:", p)

    // 3. Dereference (lấy giá trị qua con trỏ)
    fmt.Println("Giá trị tại địa chỉ p:", *p)

    // 4. Thay đổi giá trị gốc thông qua pointer
    *p = 999
    fmt.Println("Giá trị mới của a:", a)
}
```

📌 Kết quả:

* Khi thay đổi `*p` → **`a` cũng đổi theo**

---

## 4️⃣ Khi nào nên dùng Pointer?

### 🔹 A. Thay đổi dữ liệu gốc (Mutability)

Go **luôn copy giá trị** khi truyền biến vào hàm.

```go
func changeValue(x int) {
    x = 100
}
```

👉 Gọi hàm xong, biến bên ngoài **KHÔNG đổi**.

➡️ Muốn đổi thật → **truyền pointer**:

```go
func changeValue(x *int) {
    *x = 100
}
```

---

### 🔹 B. Hiệu năng (Performance)

Giả sử có struct `User` nặng **10MB**:

* ❌ Không dùng pointer → copy 10MB mỗi lần truyền
* ✅ Dùng pointer → chỉ truyền **địa chỉ (vài byte)**

👉 Pointer giúp **nhanh hơn + tiết kiệm RAM**.

---

## 5️⃣ Nil Pointer (Zero Value)

Con trỏ chưa gán địa chỉ sẽ có giá trị mặc định là `nil`.

```go
var ptr *int
fmt.Println(ptr) // nil
```

⚠️ **CẢNH BÁO QUAN TRỌNG**

```go
*ptr = 10 // ❌ PANIC – chương trình crash
```

👉 Luôn kiểm tra:

```go
if ptr != nil {
    *ptr = 10
}
```

---

## 6️⃣ Bài tập thực hành (Exercises)

### 🧩 Bài 1: Basic Pointer

* Khai báo biến `number`
* Tạo pointer trỏ tới `number`
* Thay đổi giá trị của `number` thông qua pointer

---

### 🧩 Bài 2: Swap Function (Hoán đổi)

Viết hàm:

```go
func swap(a *int, b *int)
```

Input:

```go
x := 10
y := 20
swap(&x, &y)
```

Output:

```text
x = 20
y = 10
```

---

### 🧩 Bài 3: Update Struct bằng Pointer

* Tạo struct `User { Name string, Age int }`
* Viết hàm `UpdateAge(user *User, newAge int)`
* Gọi hàm và kiểm tra giá trị `Age` có thay đổi không

---

## 🔑 Tổng kết nhanh

* Go **không có reference**, chỉ có **value & pointer**
* Pointer giúp:

    * Thay đổi dữ liệu gốc
    * Tối ưu hiệu năng
* Tránh dùng pointer khi không cần (giữ code đơn giản)

---

🚀 **Nắm chắc pointer là bạn đã vượt qua một cửa ải lớn của Go rồi đó!**
Happy Coding 💙
