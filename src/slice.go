package main

import "fmt"

func sum(s ...int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func slice() {
	sl := []int{1, 2, 3}
	sl = append(sl, 4)

	sl2 := make([]int, len(sl))

	copy(sl2, sl)
	sl2[0] = 2
	fmt.Println(sl, sl2)

	fmt.Println(sum(sl...))
}
