package binarySearch
/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

The array may contain duplicates.

Example 1:

Input: [1,3,5]
Output: 1
Example 2:

Input: [2,2,2,0,1]
Output: 0
 */
func FindMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	lo, hi := 0, len(nums)-1
	min := nums[lo]
	for lo <= hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		if nums[m] < min {
			min = nums[m]
		}
		// Check right
		if nums[m] < nums[hi] {
			hi = m - 1
			// Check left
		} else if nums[m] > nums[hi] {
			lo = m + 1
		} else {
			// hi == nums[m] so reduce by one to check previous element
			hi -= 1
		}
	}

	return min
}
