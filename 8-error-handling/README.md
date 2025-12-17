# 8. Error Handling in Go (Xử lý lỗi)

Trong Go, **xử lý lỗi (Error Handling)** là một trong những phần *ăn tiền* nhất và cũng là điểm khác biệt lớn so với Java hay Python.

Go **không dùng `try/catch`**. Thay vào đó, **lỗi được coi là một giá trị (value)** và được xử lý một cách tường minh ngay tại nơi nó xảy ra.

---

## 1️⃣ Triết lý của Go về Lỗi

### 🔸 So sánh nhanh

* **Java / Python**

    * `throw Exception` → `try / catch`
    * Luồng chương trình có thể bị nhảy khó kiểm soát

* **Golang**

    * Hàm **trả về lỗi như một giá trị**
    * Lập trình viên **bắt buộc phải kiểm tra lỗi**

### 🔥 Ưu điểm

* Code rõ ràng, dễ đọc
* Không có lỗi bị "nuốt" ngầm
* Luồng chạy tuyến tính, dễ debug

---

## 2️⃣ Cú pháp cơ bản

Trong Go, `error` là một **interface có sẵn**.
Giá trị mặc định của `error` là `nil` (tức là **không có lỗi**).

### ✍️ Khai báo hàm có trả về lỗi

Thông thường, hàm sẽ trả về **2 giá trị**:

```go
(kết quả, error)
```

Ví dụ:

```go
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("không thể chia cho 0")
    }
    return a / b, nil
}
```

---

### ✅ Pattern kinh điển: `if err != nil`

Đây là đoạn code bạn sẽ gặp **cả đời** khi code Go 😄

```go
result, err := divide(10, 0)

if err != nil {
    fmt.Println("Lỗi rồi:", err)
    return
}

fmt.Println("Kết quả là:", result)
```

👉 Nếu code chạy được xuống dưới `if err != nil` → nghĩa là **KHÔNG có lỗi**.

---

## 3️⃣ Tạo lỗi (Custom Errors)

### 🔹 Cách 1: `errors.New()`

Dùng cho lỗi **đơn giản**, chỉ có chuỗi text cố định.

```go
import "errors"

var errNotFound = errors.New("không tìm thấy dữ liệu")
```

---

### 🔹 Cách 2: `fmt.Errorf()`

Dùng khi cần **format lỗi** (chèn biến vào thông báo).

```go
age := -5
if age < 0 {
    return fmt.Errorf("tuổi không được âm: bạn nhập %d", age)
}
```

---

## 4️⃣ Panic & Recover (Nâng cao ⚠️)

### 🔥 `panic`

* Dừng chương trình **ngay lập tức**
* Dùng cho lỗi **không thể cứu vãn**

```go
panic("Lỗi nghiêm trọng!")
```

### 🛟 `recover`

* Dùng để **cứu chương trình khỏi panic**
* Thường đặt trong `defer`

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

### ⚠️ Lưu ý quan trọng

> Go **không khuyến khích** dùng `panic` cho logic thông thường.
> 👉 Hãy **ưu tiên return error**.

---

## 5️⃣ Bài tập thực hành

* Viết hàm chia 2 số, bắt lỗi chia cho 0
* Viết hàm kiểm tra tuổi (tuổi < 18 → trả về lỗi)
* In lỗi ra console

---

## 💡 Bài tập nâng cao (Code luôn cho nóng 🔥)

### 🎯 Yêu cầu

Tạo file `main.go`, viết hàm:

```go
CheckLogin(username string, password string) error
```

### 📌 Điều kiện

1. `username` rỗng → trả về lỗi **"Username không được để trống"**
2. `password` < 6 ký tự → trả về lỗi **"Mật khẩu quá yếu"**
3. Hợp lệ → trả về `nil`

### ▶️ Trong `main`

* Gọi hàm `CheckLogin`
* In ra:

    * `Đăng nhập thành công`
    * hoặc `Đăng nhập thất bại: <lý do lỗi>`

---

🚀 **Làm xong bài này là bạn đã nắm được tư duy Error Handling chuẩn Go rồi đó!**

Happy Coding 💙
