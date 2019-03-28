package binarySearch

/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
 */

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 == 0 {
		return median(nums2)
	}
	if n2 == 0 {
		return median(nums1)
	}
	if nums2[0] > nums1[n1-1] {
		if n1 == n2 {
			return avg(float64(nums2[0]), float64(nums1[n1-1]))
		}
		return median(append(nums1, nums2...))
	}
	if nums1[0] > nums2[n2-1] {
		if n1 == n2 {
			return avg(float64(nums1[0]), float64(nums2[n2-1]))
		}
		return median(append(nums2, nums1...))
	}

	if n1 < n2 {
		for _, v := range nums1 {
			index := bisectLeft(nums2, v)
			nums2 = append(nums2[:index], append([]int{v}, nums2[index:]...)...)

		}
		return median(nums2)
	} else {
		for _, v := range nums2 {
			index := bisectLeft(nums1, v)
			nums1 = append(nums1[:index], append([]int{v}, nums1[index:]...)...)
		}
		return median(nums1)
	}
	return -1
}

func median(nums []int) float64  {
	n := len(nums)
	m := int(uint(n) >> 1)
	if n & 1 == 0 {
		return float64(nums[m] + nums[m-1]) / 2.0
	} else {
		return float64(nums[m])
	}
}

func avg(a,b float64) float64  {
	return (a + b) / 2.0
}

func bisectLeft(a []int, target int) int {
	lo, hi := 0, len(a)
	for lo < hi {
		m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
		if a[m] < target {
			lo = m + 1
		} else {
			hi = m
		}
	}
	return lo
}
