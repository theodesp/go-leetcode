package binarySearch

/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

You may assume no duplicate exists in the array.

Example 1:

Input: [3,4,5,1,2]
Output: 1
Example 2:

Input: [4,5,6,7,0,1,2]
Output: 0
 */

func FindMinInRotatedArray(nums []int) int {
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
			return min
		}
	}

	return min
}
