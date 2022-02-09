package algo

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := [][]int{
		{5, 2, 6, 12, 9, 11, 3},
		{54, 32, 26, 12, 39, 11, 3},
	}

	for i, test := range tests {
		ShellSort(test)
		fmt.Println("test ", i, ":", test, IsSorted(test))
	}
}
