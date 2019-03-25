package binarySearch

/**
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
 */


func TwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1,-1}
	}

	lo, hi := 0, len(nums) - 1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target  {
			return []int{lo+1, hi+1}
		} else if sum > target {
			hi -= 1
		} else {
			lo += 1
		}
	}

	return []int{-1,-1}
}