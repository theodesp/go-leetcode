package binarySearch

import "sort"

/**
Find the Duplicate Number

Given an array nums containing n + 1 integers where each integer is
between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Example 1:

Input: [1,3,4,2,2]
Output: 2
Example 2:

Input: [3,1,3,4,2]
Output: 3
Note:

You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
 */

func FindDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sort.Ints(nums)
	n := len(nums)
	idx, curr := 1, nums[0]
	for idx < n {
		if curr == nums[idx] {
			return curr
		} else {
			curr = nums[idx]
		}
		idx += 1
	}
	return -1
}
