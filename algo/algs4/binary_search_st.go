package algs4

import "fmt"

type BinarySearchST struct {
	keys []Key
	vals []interface{}
	n    int
}

func NewBinarySearchST(capacity int) *BinarySearchST {
	return &BinarySearchST{
		keys: make([]Key, capacity),
		vals: make([]interface{}, capacity),
	}
}

func (bst *BinarySearchST) Get(key Key) interface{} {
	i := bst.Rank(key)
	if i < bst.n && key.compareTo(bst.keys[i]) == 0 {
		return bst.vals[i]
	}
	return nil
}

func (bst *BinarySearchST) Put(key Key, val interface{}) {
	pos := bst.Rank(key)
	fmt.Println("pos", pos)
	if pos < bst.n && key.compareTo(bst.keys[pos]) == 0 {
		bst.vals[pos] = val
		return
	}
	// resize
	if bst.n == len(bst.keys) {
		bst.resize(2 * len(bst.keys))
	}
	for i := bst.n; i > pos; i-- {
		bst.keys[i] = bst.keys[i-1]
		bst.vals[i] = bst.vals[i-1]
	}
	bst.keys[pos] = key
	bst.vals[pos] = val
	bst.n++
}

func (bst *BinarySearchST) Rank(key Key) int {
	left, right := 0, bst.n-1
	for left <= right {
		mid := (right-left)/2 + left
		cmp := key.compareTo(bst.keys[mid])
		if cmp < 0 {
			right = mid - 1
		} else if cmp > 0 {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func (bst *BinarySearchST) resize(capacity int) {
	newKeys := make([]Key, capacity)
	copy(newKeys, bst.keys)
	bst.keys = newKeys

	newVals := make([]interface{}, capacity)
	copy(newVals, bst.vals)
	bst.vals = newVals
}

func (bst *BinarySearchST) Keys() []Key {
	return bst.keys
}

func (bst *BinarySearchST) Size() int {
	return bst.n
}
