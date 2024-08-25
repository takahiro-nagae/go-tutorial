package main

import "fmt"

func outer() {
	var outerVar string = "outer"
	fmt.Println(outerVar)
}

func varPrint() {
	var i int = 100
	fmt.Println(i)

	var s string = "Hello, Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Print(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	// 暗黙的な型定義
	// 関数内でしか定義できない
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	outer()
}
