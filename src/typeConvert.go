package main

import (
	"fmt"
	"strconv"
)

func typeConvert() {
	var s string = "100"
	fmt.Printf("%T\n", s)

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
	fmt.Println(i)
}
