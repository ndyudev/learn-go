package main

import "fmt"

type UserDAO interface {
	create() string
	detele() string
	update() string
	findById() string
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

func main() {
	u2 := User{}
	u1 := User{}

	fmt.Println(u2.create())
	fmt.Println(u1.create())
}
