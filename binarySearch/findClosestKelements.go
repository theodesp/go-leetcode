package binarySearch

/**
Find K Closest Elements
  Go to Discuss
Given a sorted array, two integers k and x, find the k closest elements to x in the array. The result should also be sorted in ascending order. If there is a tie, the smaller elements are always preferred.

Example 1:
Input: [1,2,3,4,5], k=4, x=3
Output: [1,2,3,4]
Example 2:
Input: [1,2,3,4,5], k=4, x=-1
Output: [1,2,3,4]
Note:
The value k is positive and will always be smaller than the length of the sorted array.
Length of the given array is positive and will not exceed 104
Absolute value of elements in the array and x will not exceed 104

 */

func FindClosestElements(arr []int, k int, x int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if x < arr[0] {
		return arr[:k]
	}
	if x > arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}

	lo, hi := 0, len(arr) - 1
	for lo < hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		if arr[m] > x {
			hi = m - 1
		} else if arr[m] < x {
			lo = m + 1
		} else {
			lo = m
			break
		}
	}
	left := lo
	right := lo+1
	for k > 0 {
		if right >= len(arr) || (left >= 0 && x-arr[left] <= arr[right] - x) {
			left -= 1
		} else {
			right += 1
		}

		k -= 1
	}

	return arr[left+1:right]
}