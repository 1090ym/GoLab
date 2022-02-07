package bynarysearch

func search(nums []int, target int) bool {
	var (
		l = 0
		r = len(nums) - 1
	)
	for l <= r {
		mid := (l-r)/2 + r
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}
