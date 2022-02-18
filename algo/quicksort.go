package algo

func QuickSort(a []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	m := partition(a, lo, hi)
	QuickSort(a, lo, m)
	QuickSort(a, m+1, hi)
	return
}

func partition(a []int, lo int, hi int) int {
	i, j := lo+1, hi
	for {
		for i <= j && less(a, i, lo) {
			i++
		}
		for i <= j && less(a, lo, j) {
			j--
		}
		if i >= j {
			break
		}
		exch(a, i, j)
	}
	exch(a, lo, j)
	return j
}
