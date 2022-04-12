package cmd

import (
	"GoLab/algo/algs4"
	"fmt"
	"testing"
)

func TestBinarySearchST(t *testing.T) {
	bst := algs4.NewBinarySearchST(5)
	bst.Put(algs4.StringKey("aa"), 100)
	bst.Put(algs4.StringKey("cc"), 100)
	bst.Put(algs4.StringKey("gg"), 100)
	bst.Put(algs4.StringKey("ff"), 100)
	bst.Put(algs4.StringKey("pp"), 200)
	fmt.Println("keys", bst.Keys())
	fmt.Println("get pp", bst.Get(algs4.StringKey("pp")))
	bst.Put(algs4.StringKey("pp"), 300)
	fmt.Println("get pp", bst.Get(algs4.StringKey("pp")))
}
