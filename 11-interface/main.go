package main

import "fmt"

type UserDAO interface {
	create() string
	detele() string
	update() string
	findById() string
	read()
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
	u1 := User{}

	u2.read()
	u1.read()

}
