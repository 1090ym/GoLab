package algs4

func InsertSort(nums []int) {
	N := len(nums)
	for i := 1; i < N; i++ {
		for j := i; j > 0 && less(nums, j, j-1); j-- {
			exch(nums, j, j-1)
		}
	}
}
