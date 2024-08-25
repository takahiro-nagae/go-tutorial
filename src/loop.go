package main

func loop() {

	sl := []string{"hoge", "fuga", "piyo"}
	m := map[string]int{"hoge": 1, "fuga": 2, "piyo": 3}

	for i, v := range sl {
		println(i, v)
	}

	for _, v := range sl {
		println(v)
	}

	for k, v := range m {
		println(k, v)
	}
}
