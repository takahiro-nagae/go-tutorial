package main

import "fmt"

func switchType() {
	var x interface{}
	x = 1

	i, isInt := x.(int)
	fmt.Println(i, isInt)

	switch x.(type) {
	case int:
		println("int")
	case string:
		println("string")
	default:
		println("unknown")
	}

	switch v := x.(type) {
	case int:
		println(v, "int")
	case string:
		println(v, "string")
	default:
		println(v, "unknown")
	}
}
