package main

import (
	"fmt"
	"sync"
	"testing"
)

var m sync.Map

func TestSync(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			m.Store(n, n)
			value, _ := m.Load(n)
			fmt.Println("key:", n, ", value:", value)
			wg.Done()
		}(i) // 需要传递参数，使用闭包会出现错误
	}
	wg.Wait()
}
