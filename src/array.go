package main

import "fmt"

func array() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B", "C"}
	fmt.Println(arr2)

	fmt.Println(arr2[0])
	fmt.Println(len(arr2))
}
