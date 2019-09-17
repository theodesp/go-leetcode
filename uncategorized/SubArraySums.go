package uncategorized

/*
Given an array of integers and an integer k,
you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
Note:
The length of the array is in range [1, 20,000].
The range of numbers in the array is [-1000, 1000]
and the range of the integer k is [-1e7, 1e7].

Example explanation: [1,1] and [1,1]
 */

/* Solution 1: Brute force
 Two for loops first i from 0 to len(arr). Second from i+1 to len(arr).
Check if current sum equals k. If yes then increase the count.

Return count

Time complexity : Considering every possible subarray takes O(n^2)
 ) time. Finding out the sum of any subarray takes O(1)
time after the initial processing of O(n) for creating the cumulative sum array.
Space complexity : O(1). Constant space is used.
 */

func SubarraySumBrute(nums []int, k int) int {
	count := 0
	for i := 0;i < len(nums); i += 1 {
		curr := 0
		for j := i;j < len(nums); j += 1 {
			curr += nums[j]
			if curr == k {
				count += 1
			}
		}
	}
	return count
}

/* Solution 2: Using Cache
 we make use of a hashmap mapmap which is used to store the cumulative sum upto
all the indices possible along with the number of times the same sum occurs.
We store the data in the form: (sum_i, no. of occurences of sum_i)(sum
i,no.ofoccurencesofsum
i). We traverse over the array nums and keep on finding the cumulative sum.
Every time we encounter a new sum, we make a new entry in the hashmap corresponding
to that sum. If the same sum occurs again, we increment the count corresponding
to that sum in the hashmap. Further, for every sum encountered, we also
determine the number of times the sum sum-ksum−k has occured already, since
it will determine the number of times a subarray with sum kk has occured upto the
current index. We increment the count by the same amount.

Return count
*/

func SubarraySum(nums []int, k int) int {
	count := 0
	curr := 0
	cache := make(map[int]int)
	cache[0] = 1
	for i := 0;i < len(nums); i += 1 {
		curr += nums[i]
		// we previously store the count here for all sums from 0 to i
		if val, ok := cache[curr - k]; ok {
			count += val
		}
		cache[curr] = cache[curr] + 1
	}
	return count
}

