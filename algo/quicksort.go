package algo

func partition(a []int, lo int, hi int) int {
	i, j := lo+1, hi
	for {
		for less(a, i, lo) && i <= j {
			i++
		}
		for less(a, lo, j) && i <= j {
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
