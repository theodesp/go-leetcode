package binarySearch

import "math"

/*
Find a peak element in a 2D array
 */

/*Solution:
Consider mid column and find maximum element in it.
Let index of mid column be ‘mid’, value of maximum element in mid column be ‘max’ and maximum element be at ‘mat[max_index][mid]’.
If max >= A[index][mid-1] & max >= A[index][pick+1], max is a peak, return max.
If max < mat[max_index][mid-1], recur for left half of matrix.
If max < mat[max_index][mid+1], recur for right half of matrix.
 */
func FindPeakElement2d(nums [][]int) int {
	return findPeakElementRecur(nums, len(nums) >> 1)
}

func findPeakElementRecur(nums [][]int, m int) int  {
	i, max := findMax(nums, m)
	// If we are on the first or last column,
	// max is a peak
	if m == 0 || m == len(nums) - 1 {
		return max
	}
	// If max is less than its left
	if max < nums[i][m-1] {
		return findPeakElementRecur(nums, m - int(math.Ceil(float64(m >> 1))))
	}
	// if max is less than its right
	if max < nums[i][m+1] {
		return findPeakElementRecur(nums, m + int(math.Ceil(float64(m >> 1))))
	}
	// If mid column maximum is also peak
	return max
}

func findMax(nums [][]int, m int) (int, int) {
	var curMax int
	var index int
	for i, _ := range nums {
		if nums[i][m] > curMax {
			curMax = nums[i][m]
			index = i
		}
	}
	return index, curMax
}