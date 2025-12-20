# 10. Structs in Go (Xương sống của Go 🏗️)

Nếu bạn đã từng code **Java**, hãy nhớ nhanh một câu này cho dễ hình dung:

> ❌ Go **KHÔNG có Class**
> ✅ **Struct chính là thứ thay thế Class trong Go**

Struct dùng để **gom nhiều dữ liệu khác nhau** (int, string, slice, struct khác, …) vào chung một đối tượng.

---

## 1️⃣ Khai báo Struct (Blueprint)

Dùng từ khóa `type` + `struct` để định nghĩa một cái "khuôn".

```go
package main

import "fmt"

// Định nghĩa Struct
// Tên viết HOA chữ cái đầu -> Public (package khác dùng được)
type Student struct {
    ID      int
    Name    string
    Age     int
    Classes []string // Struct có thể chứa slice
}

func main() {
    // --- Cách 1: Khai báo đầy đủ (Khuyên dùng) ---
    student1 := Student{
        ID:   1,
        Name: "Duy Nguyen",
        Age:  20,
        Classes: []string{"Math", "Code"},
    }

    // --- Cách 2: Zero Value ---
    var student2 Student
    // ID = 0, Name = "", Age = 0, Classes = nil

    student2.Name = "Nam"
    student2.Age = 18

    fmt.Println("Student 1:", student1)
    fmt.Println("Student 2:", student2)

    // Truy cập field
    fmt.Println("Tên SV1:", student1.Name)
}
```

📌 **Zero Value** rất quan trọng trong Go → giúp code an toàn, ít null bug.

---

## 2️⃣ Struct & Pointer (CỰC KỲ QUAN TRỌNG ⭐️)

Struct ngoài đời thường **khá nặng** (nhiều field).
Nếu truyền struct bình thường → Go sẽ **COPY toàn bộ struct**.

👉 Thực tế **99% dùng Pointer to Struct**.

```go
ptrStudent := &Student{
    ID:   2,
    Name: "Huy",
    Age:  22,
}

// Go tự động dereference
fmt.Println(ptrStudent.Name)
```

👉 Bình thường phải viết `(*ptrStudent).Name` nhưng Go cho viết gọn luôn.

---

## 3️⃣ Methods – Hàm gắn với Struct

Trong Java:

```java
class Student {
    void study() {}
}
```

Trong Go:
➡️ Hàm viết **bên ngoài**, gắn vào struct bằng **Receiver**.

---

## 4️⃣ Value Receiver vs Pointer Receiver

### 🔹 A. Value Receiver (Bản COPY)

Dùng khi:

* Chỉ đọc dữ liệu
* Không cần sửa struct

```go
func (s Student) DisplayInfo() {
    fmt.Println("Tên:", s.Name)
}
```

⚠️ Thay đổi trong hàm **KHÔNG ảnh hưởng bản gốc**.

---

### 🔹 B. Pointer Receiver (Bản GỐC ⭐️)

Dùng khi:

* Cần sửa dữ liệu struct
* Struct lớn (tối ưu RAM)

```go
func (s *Student) Birthday() {
    s.Age++
}
```

👉 **Best Practice:**

> Khi phân vân → **dùng Pointer Receiver**.

---

## 5️⃣ Ví dụ tổng hợp (Thực chiến)

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

// Value Receiver (Fake)
func (u User) changeNameFake(newName string) {
    u.Name = newName
    fmt.Println("Trong hàm fake:", u.Name)
}

// Pointer Receiver (Real)
func (u *User) changeNameReal(newName string) {
    u.Name = newName
    fmt.Println("Trong hàm real:", u.Name)
}

func main() {
    user := User{Name: "Duy", Age: 20}

    fmt.Println("--- Value Receiver ---")
    user.changeNameFake("Hacker")
    fmt.Println("Sau fake:", user.Name)

    fmt.Println("--- Pointer Receiver ---")
    user.changeNameReal("Nhat Duy")
    fmt.Println("Sau real:", user.Name)
}
```

---

## 6️⃣ Bài tập thực hành 🧩

### 🧠 Bài 1: Rectangle

```go
type Rectangle struct {
    Width  int
    Height int
}
```

* Viết method `Area()` → **Value Receiver** ❓
* Viết method `Scale(factor int)` → **Pointer Receiver** ❓

👉 Tự trả lời:

* Area chỉ đọc dữ liệu → Value
* Scale thay đổi kích thước → Pointer

---

## 7️⃣ Tổng kết nhanh

* Struct = Class (phiên bản Go)
* Không có constructor, getter, setter rườm rà
* Pointer + Struct = combo mạnh nhất của Go
* Hiểu Struct + Pointer → học Go nhanh gấp đôi 🚀

---

💙 **Nắm chắc bài này là bạn đã bước vào "core Go" thật sự rồi đó!**
Happy Coding!
