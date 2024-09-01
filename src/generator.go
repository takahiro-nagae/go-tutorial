package main

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func generator() {
	ints := integers()
	println(ints())
	println(ints())
	println(ints())
	println(ints())
	println(ints())
}
