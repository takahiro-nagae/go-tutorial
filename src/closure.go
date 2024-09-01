package main

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func closure() {
	f := later()
	println(f("Golang"))
	println(f("is"))
	println(f("awesome"))
	println(f("!!!"))
}
