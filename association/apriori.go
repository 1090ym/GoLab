package association

import (
	"sort"
)

type itemset struct {
	items   []int
	count   int
	support float64
}

// ScanC1 生成1项的所有项集
func ScanC1(data [][]int) []itemset {
	m := make(map[int]int)
	//count := 0
	for _, trans := range data {
		for _, item := range trans {

			m[item]++

		}
	}

	res := make([]itemset, 0, len(m))
	for k, v := range m {
		itemSet := itemset{
			items: []int{k},
			count: v,
		}
		res = append(res, itemSet)
	}
	return res
}

// ScanFreqItemSet 统计Ck中的频繁项集
func ScanFreqItemSet(data [][]int, Ck []itemset, minSupport float64) []itemset {
	ssCnt := make([]itemset, 0, len(Ck))
	for _, c := range Ck {
		set := itemset{}
		for _, trans := range data {
			// 如果item-set是transaction的一个子集，则进行count
			if IsSubset(c.items, trans) {
				set.items = c.items
				set.count++
			}
		}
		//fmt.Println("support:", float64(set.count/len(data)), " data len:", len(data))
		if float64(set.count)/float64(len(data)) >= minSupport {
			set.support = float64(set.count) / float64(len(data))
			ssCnt = append(ssCnt, set)
		}
	}
	return ssCnt
}

// AprioriGen 根据k项的频繁项集生成k+1项的所有项集
func AprioriGen(Ck []itemset) []itemset {
	k := len(Ck)
	res := make([]itemset, 0)
	for i := 0; i < k-1; i++ {
		for j := i + 1; j < k; j++ {
			set, ok := mergeItemSet(Ck[i], Ck[j])
			if ok {
				res = append(res, set)
			}
		}
	}
	return res
}

// Apriori Apriori算法，收集频繁项集
func Apriori(dataset [][]int, minSupport float64) [][]itemset {
	C1 := ScanC1(dataset)
	L := make([][]itemset, 0)
	L1 := ScanFreqItemSet(dataset, C1, minSupport)
	L = append(L, L1)
	k := 2
	for len(L[k-2]) > 0 {
		Ck := AprioriGen(L[k-2])
		Lk := ScanFreqItemSet(dataset, Ck, minSupport)
		if len(Lk) == 0 {
			break
		}
		L = append(L, Lk)
		k++
	}
	return L
}

// mergeItemSet 合并2个itemset
func mergeItemSet(setA itemset, setB itemset) (itemset, bool) {
	mergeSet := itemset{}
	if len(setA.items) != len(setB.items) {
		return mergeSet, false
	}
	k := len(setA.items)
	flag := true
	sort.Ints(setA.items)
	sort.Ints(setB.items)
	for i := 0; i < k-1; i++ {
		if setA.items[i] != setB.items[i] {
			flag = false
		}
	}
	if flag {
		mergeSet.items = append(setA.items[:k-1], setA.items[k-1])
		mergeSet.items = append(mergeSet.items, setB.items[k-1])
	}
	return mergeSet, flag
}

// IsSubset 判断setA是否是setB的子集
func IsSubset(setA []int, setB []int) bool {
	for _, a := range setA {
		flag := false
		for _, b := range setB {
			if a == b {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
