package cmd

import (
	"GoLab/algo/algs4"
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
		algs4.QuickSort(test, 0, len(test)-1)
		fmt.Println("test ", i, ":", test, algs4.IsSorted(test))
	}
}
