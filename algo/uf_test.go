package algo

import (
	"fmt"
	"testing"
)

func TestUf(t *testing.T) {
	uf := UF{}
	uf.NewUF(10)
	nums := []struct {
		p int
		q int
	}{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{5, 0},
		{7, 2},
		{6, 1},
	}

	for _, num := range nums {
		if uf.Connected(num.p, num.q) {
			continue
		}
		uf.Union(num.p, num.q)
	}
	fmt.Println("uf count:", uf.Count())
	fmt.Println("uf id:", uf.Id())
}
