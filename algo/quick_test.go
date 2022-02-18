package algo

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	a := []int{15, 2, 6, 12, 9, 11, 3}
	b := []int{54, 32, 26, 12, 39, 11, 3}
	fmt.Println(partition(a, 0, len(a)-1))
	fmt.Println(a)
	fmt.Println(partition(b, 0, len(b)-1))
	fmt.Println(b)

}
