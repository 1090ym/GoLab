package algs4

func SelectionSort(nums []int) {
	N := len(nums)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if less(nums, j, min) {
				min = j
			}
		}
		exch(nums, i, min)
	}
}
