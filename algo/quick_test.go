package algo

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	a := []int{5, 2, 6, 12, 9, 11, 3}
	fmt.Println(partition(a, 0, len(a)-1))
	fmt.Println(a)
}
