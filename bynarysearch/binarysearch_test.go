package bynarysearch

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		res    bool
	}{
		{
			nums:   []int{1, 3, 5, 7, 8, 9},
			target: 7,
			res:    true,
		},
		{
			nums:   []int{2, 5, 5, 7, 9, 11},
			target: 9,
			res:    true,
		},
	}

	for _, test := range tests {
		res := search(test.nums, test.target)
		assert.Equal(t, res, test.res)
	}
}
