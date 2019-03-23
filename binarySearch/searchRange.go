package binarySearch

/**
Given an array of integers nums sorted in ascending order,
find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
 */

func SearchRange(nums []int, target int) []int {
	l := len(nums)
	if l == 0 {
		return []int{-1, -1}
	}
	if l == 1 && nums[0] != target {
		return []int{-1, -1}
	}

	if l == 1 && nums[0] == target {
		return []int{0, 0}
	}

	lo, hi, tip := 0, l, -1
	for lo < hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		if nums[m] < target {
			lo = m + 1
		} else if nums[m] > target {
			hi = m
		} else {
			tip = m
			break
		}
	}
	if tip == -1 {
		return []int{-1, -1}
	}
	lo, hi = tip, tip
	for i:= tip-1; i >= 0; i-=1 {
		if nums[i] != target {
			break
		}
		lo = i
	}
	for i:= tip+1; i < l; i+=1 {
		if nums[i] != target {
			break
		}
		hi = i
	}
	return []int{lo, hi}
}