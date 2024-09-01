package main

import "fmt"

func mapSample() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
