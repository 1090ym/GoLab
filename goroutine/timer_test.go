package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// 1.timer基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	// 2s后会从channel中接收到信息
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	//timer2 := time.NewTimer(2 * time.Second)
	t1 = time.Now()
	fmt.Printf("t1:%v\n", t1)
	<-time.After(2 * time.Second)
	fmt.Println("延迟2s：", time.Now())

	// 4.停止定时器
	timer3 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer3.C
		fmt.Println("定时器执行了")
	}()
	b := timer3.Stop()
	if b {
		fmt.Println("timer3已经关闭")
	}

	ticker := time.NewTicker(time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		i := 0
		for {
			i++
			fmt.Println("ticker:", <-ticker.C)
			if i == 4 {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
