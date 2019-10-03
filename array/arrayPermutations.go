package array

/*
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

Solution using recursion:
First have a function called permutate that accepts an array and 2 indexes start and end.
The base case is when start == end then we just return the array.
In any other case we iterate from start to end and each time we swap the start with i we
call recursively permutate and restore the swap:

[1,2,3] -> start=0, end=2 keeps 1 fixed
[1,2,3] -> start=1, end=2 keeps 2 fixed
[1,2,3] -> start=2, end=2 return [1,2,3]
[1,2,3] -> restore 2nd swap
[1,3,2] -> swap 3 for 2
[1,3,2] -> start=2, end=2 return [1,3,2]
[1,2,3] -> restore 2nd swap
[1,2,3] -> restore 1nd swap
[2,1,3] -> swap 2 for 1
...
 */

func Permute(nums []int) [][]int {
	result := [][]int{}
	permutateHelper(nums, 0, len(nums), func(arr []int) {
		result = append(result, arr)
	})
	return result
}

func permutateHelper(arr []int, start, end int, cb func(arr []int)) {
	if start == end {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		cb(tmp)
	} else {
		for i := start;i<end; i+=1 {
			// Keep current number fixed
			swap(arr, start, i)
			// Perform recursively
			permutateHelper(arr, start+1, end, cb)
			// Restore swap
			swap(arr, start, i)
		}
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}