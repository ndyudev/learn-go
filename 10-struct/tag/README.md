# 11. Struct Tags in Go (Metadata – Thứ "ăn tiền" trong Go 🚀)

**Struct Tags** là thứ mà **100% dự án Go thực tế đều dùng**, đặc biệt khi làm:

* Web API (JSON / REST)
* Database (ORM như GORM)
* Validation, Config, Binding dữ liệu

Nếu ví **Struct là khung sườn ngôi nhà** 🏗️
thì **Struct Tag chính là mấy tờ sticky note** 📝 dán lên từng phòng để hướng dẫn:

> “Phòng này dùng làm gì? Xuất ra sao? Lưu DB thế nào?”

---

## 1️⃣ Struct Tag là gì?

**Struct Tag** là một đoạn text nằm trong **dấu nháy ngược** (backticks `...`) ở cuối mỗi field.

```go
type User struct {
    Name string `json:"name"`
}
```

### 🔎 Bản chất

* Struct Tag là **metadata (dữ liệu mô tả dữ liệu)**
* Go compiler **KHÔNG quan tâm** tới tag
* Các thư viện bên ngoài (`encoding/json`, GORM, validator, …) sẽ **đọc tag để xử lý**

👉 Struct Tag = *hợp đồng ngầm* giữa struct của bạn và thư viện.

---

## 2️⃣ Cú pháp chuẩn (RẤT QUAN TRỌNG ⚠️)

```go
type User struct {
    //  Field      Type      Struct Tag
    //    ↓          ↓           ↓
    FirstName string `json:"first_name" db:"name"`
}
```

### 🚫 Luật bất di bất dịch

Trong `key:"value"`:

* ❌ **KHÔNG ĐƯỢC có khoảng trắng**

| Đúng          | Sai            |
| ------------- | -------------- |
| `json:"name"` | `json: "name"` |

👉 Sai là thư viện **đọc không ra tag → bug ngầm rất khó tìm**.

---

## 3️⃣ Ứng dụng phổ biến nhất: JSON (Web API)

### ❓ Vấn đề thực tế

* Field Go **phải viết HOA** để public
* JSON API thường dùng **snake_case**

```go
FullName  ❌  →  full_name  ✅
```

### ✅ Giải pháp

Dùng Struct Tag để **map tên field**.

---

## 4️⃣ Các JSON Tag hay dùng nhất

| Tag           | Ý nghĩa                     | Ví dụ                            |
| ------------- | --------------------------- | -------------------------------- |
| `json:"name"` | Đổi tên field khi xuất JSON | `Name string `json:"username"``  |
| `json:"-"`    | Ẩn hoàn toàn field          | `Password string `json:"-"``     |
| `omitempty`   | Giá trị zero → Ẩn field     | `Age int `json:"age,omitempty"`` |

### 📌 Zero value là gì?

* `int` → `0`
* `string` → `""`
* `bool` → `false`
* `slice/map` → `nil`

---

## 5️⃣ Code thực chiến: JSON Marshal (Copy chạy thử)

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    Name      string  `json:"product_name"`
    Price     float64 `json:"price"`
    IsOnSale  bool    `json:"is_on_sale,omitempty"`
    SecretKey string  `json:"-"`
}

func main() {
    p1 := Product{
        Name:      "Iphone 15",
        Price:     1000,
        IsOnSale:  true,
        SecretKey: "ABC_XYZ",
    }

    json1, _ := json.MarshalIndent(p1, "", "  ")
    fmt.Println("--- Product 1 ---")
    fmt.Println(string(json1))

    p2 := Product{
        Name:      "Samsung S24",
        Price:     900,
        IsOnSale:  false,
        SecretKey: "SECRET",
    }

    json2, _ := json.MarshalIndent(p2, "", "  ")
    fmt.Println("\n--- Product 2 ---")
    fmt.Println(string(json2))
}
```

### 🔥 Kết quả

* `product_name` được đổi tên
* `is_on_sale` **biến mất** khi `false`
* `SecretKey` **không bao giờ lộ**

---

## 6️⃣ Struct Tag với Database (GORM)

Khi làm DB, Struct Tag dùng để map struct ↔ bảng.

```go
type User struct {
    ID    uint   `gorm:"primaryKey"`
    Email string `gorm:"unique"`
    Name  string `gorm:"column:user_name"`
}
```

👉 Struct lúc này **vừa là model Go, vừa là schema DB**.

---

## 7️⃣ Bài tập thực hành (Challenge 🧠)

### 🎯 Yêu cầu

Tạo struct `Student`:

* `FirstName` → `first_name`
* `LastName` → `last_name`
* `Age` → nếu `0` thì **ẩn**
* `Password` → **không bao giờ hiện**

```go
type Student struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Age       int    `json:"age,omitempty"`
    Password  string `json:"-"`
}
```

➡️ Tạo dữ liệu, marshal sang JSON và kiểm tra output.

---

## 🔑 Tổng kết nhanh

* Struct Tag = Metadata cho struct
* Không ảnh hưởng compiler, nhưng **ảnh hưởng toàn bộ hệ sinh thái Go**
* JSON + DB + Validation → **Struct Tag là core skill**

---

🚀 **Hiểu Struct Tag là bạn đã sẵn sàng làm Go Backend thực tế rồi đó!**
Happy Coding 💙
