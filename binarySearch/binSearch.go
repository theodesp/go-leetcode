package binarySearch

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo, hi := 0, len(nums)
	for lo < hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			lo = m + 1
		} else {
			hi = m
		}
	}

	if lo != len(nums) && nums[lo] == target {
		return lo
	}

	return -1
}
