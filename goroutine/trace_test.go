package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

func TestTrace(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		return
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	fmt.Println("trace test")

}
