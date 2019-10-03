package array

/*
Find the kth largest/smallest element in an unsorted array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.

Solution: We need to use a variation of quicksort.
First we partition the elements of the list along a pivot point.
If in the first partition point we found that the index is k then we return A[p].
If the partition point is larger than k then we recursively call the first left partition
Otherwise we recursively call the right partition
 */

func FindKthLargest(nums []int, k int) int {
	return findKthLargestHelper(nums, 0, len(nums)-1,  len(nums) - k)
}

func findKthLargestHelper(nums []int, left, right, k int) int  {
	p := partition(nums, left, right)
	if p == k {
		return nums[p]
	}
	if p > k {
		return findKthLargestHelper(nums, left, p - 1, k)
	}
	return findKthLargestHelper(nums, p + 1, right, k)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	j := left
	for i := left;i<=right;i+=1 {
		if nums[i] < pivot {
			swap(nums, i, j)
			j += 1
		}
	}
	swap(nums, right, j)
	return j
}