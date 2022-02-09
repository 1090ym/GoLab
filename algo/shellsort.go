package algo

func ShellSort(nums []int) {
	N := len(nums)
	for d := N / 2; d > 0; d /= 2 {
		for i := d; i < N; i++ {
			for j := i; j >= d && less(nums, j, j-d); j -= d {
				exch(nums, j, j-d)
			}
		}
	}
}
