package binarySearch

/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
 */

func SearchRotated(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		if nums[m] == target {
			return m
		}

		if nums[m] < nums[hi] {
			if target > nums[m] && target <= nums[hi] {
				lo = m + 1
			} else {
				hi = m - 1
			}
		} else if nums[m] > nums[hi] {
			if target < nums[m] && target >= nums[lo] {
				hi = m - 1
			} else {
				lo = m + 1
			}
		} else {
			return -1
		}
	}

	return -1
}
