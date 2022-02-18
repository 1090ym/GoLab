package algo

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := [][]int{
		{12, 2, 3, 4, 9, 11, 13},
		{2, 1, 26, 12, 39, 11, 3},
	}

	for i, test := range tests {
		//SelectionSort(test)
		//InsertSort(test)
		//ShellSort(test)
		//MergeSort(test, 0, len(test)-1)
		//MergeSortBU(test)
		QuickSort(test, 0, len(test)-1)
		fmt.Println("test ", i, ":", test, IsSorted(test))
	}
}
