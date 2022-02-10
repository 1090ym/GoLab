package association

import (
	"fmt"
	"testing"
)

func TestApriori(t *testing.T) {
	datasets := [][][]int{
		{
			{1, 3, 4},
			{2, 3, 5},
			{1, 2, 3, 5},
			{2, 5},
		},
	}
	minSupport := 0.5

	for _, dataset := range datasets {
		fmt.Println(Apriori(dataset, minSupport))
	}
}
