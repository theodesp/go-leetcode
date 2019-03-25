package binarySearch

import "sort"

/**
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Note:

Each element in the result must be unique.
The result can be in any order.

 */

func Intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	seen := make(map[int]bool)
	// Sort the smallest one
	if len(nums1) < len(nums2) {
		sort.Ints(nums1)
		// Search every element of bigger array in smaller array
		// and print the element if found
		for _, v := range nums2 {
			if Search(nums1, v) != -1 && !seen[v] {
				res = append(res, v)
				seen[v] = true
			}
		}
	} else {
		sort.Ints(nums2)
		// Search every element of bigger array in smaller array
		// and print the element if found
		for _, v := range nums1 {
			if Search(nums2, v) != -1 && !seen[v] {
				res = append(res, v)
				seen[v] = true
			}
		}
	}
	return res
}
