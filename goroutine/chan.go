package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 = make(chan int, 4)

	go func() {
		for i := 0; i < 4; i++ {
			ch1 <- i + 1
			time.Sleep(time.Second)
		}
		close(ch1)
	}()

	for i := range ch1 {
		fmt.Println("ch1:", i)
	}
}
