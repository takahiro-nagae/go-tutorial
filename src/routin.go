package main

import (
	"fmt"
	"time"
)

func sub() {
	for i := 0; i < 10; i++ {
		println("sub" + fmt.Sprint(i))
		time.Sleep(100 * time.Millisecond)
	}
}

func mainRoutine() {
	go sub()

	for i := 0; i < 10; i++ {
		println("main:" + fmt.Sprint(i))
		time.Sleep(200 * time.Millisecond)
	}
}
