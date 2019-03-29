package binarySearch

import (
	"math"
)

/*
Given an array which consists of non-negative integers and an integer m, you can split the array into m non-empty continuous subarrays. Write an algorithm to minimize the largest sum among these m subarrays.

Note:
If n is the length of array, assume the following constraints are satisfied:

1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
Examples:

Input:
nums = [7,2,5,10,8]
m = 2

Output:
18

Explanation:
There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8],
where the largest sum among the two subarrays is only 18.
 */
func SplitArray(nums []int, m int) int {
	max, sum := 0.0, 0
	for _, v := range nums {
		max = math.Max(float64(v), max)
		sum += v
	}

	if m == 1 {
		return sum
	}

	lo, hi := int(max), sum
	for lo <= hi {
		mid := int(uint64(lo+hi)) >> 1
		if checkArray(nums, m, mid) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func checkArray(nums []int, m, target int) bool  {
	count, total := 1, 0
	for _, v := range nums {
		total += v
		if total > target {
			total = v
			count += 1
			if count > m {
				return false
			}
		}
	}
	return true
}