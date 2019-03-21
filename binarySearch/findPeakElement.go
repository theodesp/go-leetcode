package binarySearch

/**
A peak element is an element that is greater than its neighbors.

Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that nums[-1] = nums[n] = -∞.

Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
 */

func FindPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	lo, hi := 1, len(nums) - 1
	if nums[0] > nums[1] {
		return 0
	}

	if nums[hi] > nums[hi-1] {
		return hi
	}
	for lo < hi {
		if nums[lo-1] < nums[lo] && nums[lo] > nums[lo+1] {
			return lo
		} else {
			lo +=1
		}
	}
	return -1
}
