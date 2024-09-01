package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Users []*User

func (u User) Greet() {
	fmt.Println("Hello, " + u.Name)
}

func (u *User) setName(name string) {
	u.Name = name
}

func NewUser(name string, age int) *User {
	return &User{name, age}
}

func typeSample() {
	u := User{
		Name: "John",
		Age:  30,
	}

	println(u.Name, u.Age)
	u.Greet()

	u.setName("Doe")

	u2 := NewUser("Jane", 20)
	fmt.Println(u2)

	users := Users{}
	users = append(users, &u)
	users = append(users, u2)

	for _, user := range users {
		fmt.Println(user)
	}

	m := map[int]User{
		1: {Name: "John", Age: 30},
		2: {Name: "Jane", Age: 20},
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
