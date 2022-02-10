package algo

import "math"

func MergeSort(a []int, l int, h int) {
	if h <= l {
		return
	}
	m := (l + h) / 2
	MergeSort(a, l, m)
	MergeSort(a, m+1, h)
	Merge(a, l, m, h)
}

func MergeSortBU(a []int) {
	N := len(a)
	for sz := 1; sz < N; sz = sz + sz {
		for l := 0; l < N-sz; l += sz + sz {
			Merge(a, l, l+sz-1, int(math.Min(float64(l+sz+sz-1), float64(N-1))))
		}
	}
}

func Merge(a []int, l int, m int, h int) {
	i, j := l, m+1
	aux := make([]int, len(a), len(a))
	for k := l; k <= h; k++ {
		if i > m {
			aux[k] = a[j]
			j++
		} else if j > h {
			aux[k] = a[i]
			i++
		} else if less(a, i, j) {
			aux[k] = a[i]
			i++
		} else {
			aux[k] = a[j]
			j++
		}
	}
	for k := l; k <= h; k++ {
		a[k] = aux[k]
	}
	return
}
