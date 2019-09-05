package uncategorized

/*
Given an array of integers, return indices of the two numbers such that
they add up to a specific target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */

 /*
 Solution 1:
 Use two for loops. First loop starts with i = 0 to nums.length.
 Second loop starts with j = i + 1 to nums.length.
 At each stage we check if nums[i] + nums[j] == target
 Return indexes if true otherwise empty

 Time Complexity: O(n^2)
 Space Complexity: O(1)
  */

func TwoSumBrute(nums []int, target int) []int {
	for i := 0; i< len(nums); i+=1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{}
}

/*
Solution 2:
Use two loops.
In the first iteration, we add each element's value and its index to the table.
Then, in the second iteration we check if each element's complement (target - nums[i]target−nums[i]) exists in the table.

 Time Complexity: O(n)
 Space Complexity: O(1)
 */
func TwoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := 0; i< len(nums); i+=1 {
		cache[nums[i]] = i
	}
	for i := 0; i< len(nums); i+=1 {
		if j, ok := cache[target-nums[i]]; ok {
			if i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}