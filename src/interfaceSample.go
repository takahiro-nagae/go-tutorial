package main

import "fmt"

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number: %s, Model: %s", c.Number, c.Model)
}

func interfaceSample() {
	vs := []Stringfy{
		&Person{Name: "John", Age: 30},
		&Car{Number: "XXX-0123", Model: "Model X"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
