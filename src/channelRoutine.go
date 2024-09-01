package main

import (
	"fmt"
	"time"
)

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func channelRoutine() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	go receiver(ch1)
	go receiver(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(100 * time.Millisecond)
		i++
	}

	close(ch1)
	close(ch2)
}
