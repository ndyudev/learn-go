# 12. Interfaces in Go (Linh hồn của Go 🚧)

**Interface** là một trong những khái niệm **quan trọng bậc nhất** của Go.
Nó giúp code:

* Linh hoạt
* Dễ mở rộng
* Dễ test (mock)

Nếu bạn đã từng học **Java**, đọc phần này là *vỡ óc theo hướng tích cực* liền 😎

---

## 1️⃣ Tư duy khác biệt: Java vs Go

### 🔸 Java – Explicit (Khai báo rõ ràng)

```java
class Dog implements Animal { }
```

👉 Bạn phải **nói thẳng**: *"Tôi implement interface này"*.

---

### 🔸 Go – Implicit (Ngầm định)

Go **KHÔNG cần** từ khóa `implements`.

> Nếu một struct **có đủ method** mà interface yêu cầu → Go tự coi như struct đó đã implement interface.

### 🦆 Duck Typing

> *Nếu nó đi như vịt, kêu như vịt, bơi như vịt → thì nó là vịt.*

Go **không quan tâm bạn là ai**, chỉ quan tâm **bạn làm được gì**.

---

## 2️⃣ Cấu trúc của Interface

Interface chỉ là một **bản hợp đồng** gồm các method.
👉 **KHÔNG chứa dữ liệu**.

```go
type DiChuyen interface {
    Run()
    Stop()
}
```

Bất kỳ struct nào có đủ `Run()` và `Stop()` → **tự động là DiChuyen**.

---

## 3️⃣ Ví dụ kinh điển: Animal Speak 🐶🐱

### 📌 Code hoàn chỉnh (Copy chạy thử)

```go
package main

import "fmt"

// 1. Interface (Hợp đồng)
type Animal interface {
    Speak() string
}

// 2. Struct (Đối tượng thực thi)
type Dog struct{}
type Cat struct{}

// 3. Implement interface (KHÔNG cần implements)
func (d Dog) Speak() string {
    return "Gâu gâu!"
}

func (c Cat) Speak() string {
    return "Meow meow!"
}

// 4. Hàm dùng interface (Polymorphism)
func MakeSound(a Animal) {
    fmt.Println("Con vật kêu:", a.Speak())
}

func main() {
    dog := Dog{}
    cat := Cat{}

    MakeSound(dog)
    MakeSound(cat)

    animals := []Animal{Dog{}, Cat{}, Dog{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

### 🔥 Điểm ăn tiền

* `MakeSound()` **không biết Dog hay Cat**
* Chỉ cần biết: *"Mày có Speak() không?"*

---

## 4️⃣ Empty Interface – `interface{}` / `any`

### ❓ Interface rỗng là gì?

```go
interface{}
```

Interface **không yêu cầu method nào** → **MỌI kiểu dữ liệu đều thỏa mãn**.

👉 Nó giống `Object` trong Java.

Từ Go 1.18+, dùng `any` cho gọn:

```go
func printAnything(v any) {
    fmt.Println(v)
}
```

### 📌 Ví dụ

```go
func main() {
    printAnything(100)
    printAnything("Hello")
    printAnything(3.14)
    printAnything(Dog{})
}
```

⚠️ Lưu ý: Dùng `any` nhiều quá → **mất type safety**.

---

## 5️⃣ Lưu ý CỰC QUAN TRỌNG về Pointer Receiver ⚠️

Nếu method được implement bằng **Pointer Receiver**, thì **interface chỉ nhận con trỏ**.

```go
func (d *Dog) Speak() string {
    return "Gâu gâu!"
}

func main() {
    var a Animal

    a = Dog{}   // ❌ LỖI
    a = &Dog{}  // ✅ ĐÚNG
}
```

👉 Nhớ rule này để tránh bug compile khó hiểu.

---

## 6️⃣ Bài tập thử thách (Challenge 🧠)

### 🎯 Yêu cầu

Tạo interface:

```go
type Geometry interface {
    Area() float64
    Perimeter() float64
}
```

Tạo struct:

* `Rectangle` → width, height
* `Circle` → radius

Viết hàm:

```go
func PrintInfo(g Geometry)
```

👉 In ra **diện tích** và **chu vi** của từng hình.

---

## 7️⃣ Tổng kết nhanh

* Interface = hợp đồng hành vi
* Go dùng **implicit interface** → cực kỳ linh hoạt
* Interface + Struct + Pointer = nền tảng Go backend
* Code càng lớn → interface càng quan trọng 🚀

---

💙 **Hiểu Interface là bạn đã chạm tới "linh hồn" của Go rồi đó!**
Happy Coding!
