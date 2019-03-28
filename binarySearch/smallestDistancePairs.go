package binarySearch

import "sort"

/**
Given an integer array, return the k-th smallest distance among all the pairs.
The distance of a pair (A, B) is defined as the absolute difference between
A and B.

Example 1:
Input:
nums = [1,3,1]
k = 1
Output: 0
Explanation:
Here are all the pairs:
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
Then the 1st smallest distance pair is (1,1), and its distance is 0.
Note:
2 <= len(nums) <= 10000.
0 <= nums[i] < 1000000.
1 <= k <= len(nums) * (len(nums) - 1) / 2.
 */

func SmallestDistancePair(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	sort.Ints(nums)
	lo, hi := 0, nums[len(nums)-1] - nums[0]
	for lo < hi {
		m := int(uint(lo+hi) >> 1)
		if countPairs(nums, m) >= k {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return lo
}

func countPairs(nums []int, k int) int  {
	n, count := len(nums), 0

	for r, c := 0, 1;r < n-1;r+=1 {
		for c < n && (nums[c] - nums[r]) <= k {
			c += 1
		}
		count += c-r-1
	}
	return count
}