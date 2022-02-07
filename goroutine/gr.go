package main

import (
	"fmt"
	"time"
)

func gr() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("go routine:", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main:", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
