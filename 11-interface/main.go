package main

import (
	"fmt"
	"learn-go/11-interface/cat"
	dog "learn-go/11-interface/dog"
)

func PrintValue(value any) {
	fmt.Println("Value:", value)
}

type UserDAO interface {
	create() string
	detele() string
	update() string
	findById() string
	read()
}

type Animal interface {
	Speak() string
	GetName() string
}

func MakeSound(a Animal) {
	fmt.Println(a.GetName())
	fmt.Println(a.Speak())
}

type User struct {
	id   string
	name string
	age  int
}

func (u User) create() string {
	fmt.Println(u.id)
	fmt.Println(u.name)
	fmt.Println(u.age)
	return u.id
}

func (u User) read() {
	fmt.Println(u.id)
	fmt.Println(u.name)
	fmt.Println(u.age)
}

func main() {
	u2 := User{
		id:   "1",
		name: "Dyu",
		age:  20,
	}
	u1 := User{
		id:   "2",
		name: "trh",
		age:  19,
	}
	u3 := User{
		id:   "3",
		name: "huyen tra",
		age:  19,
	}

	dogPet := dog.New("Bully")

	catPet := cat.NewCat("")
	fmt.Println(dogPet.Speak())
	fmt.Println(catPet.Speak())
	u2.read()
	u1.read()
	u3.read()

	MakeSound(dogPet)

	PrintValue("2")
	var name any = 10
	number := name
	fmt.Println(number)
}
