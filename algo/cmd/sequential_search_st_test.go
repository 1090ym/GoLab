package cmd

import (
	"GoLab/algo/algs4"
	"fmt"
	"testing"
)

func TestSequentialSearchST(t *testing.T) {
	st := algs4.NewSequentialSearchST()
	st.Put(0, "name")
	st.Put(2, "city")
	st.Put(4, "count")
	st.Put(4, "user")
	st.Put(8, "work")

	fmt.Println("0", st.Get(0))
	fmt.Println("8", st.Get(8))
	fmt.Println("2", st.Get(2))
	fmt.Println("4", st.Get(4))
}
