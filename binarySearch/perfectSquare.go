package binarySearch

/**
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

Input: 16
Output: true
Example 2:

Input: 14
Output: false
 */

func IsPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	if num == 2 {
		return false
	}
	var i = 2
	for ;i * i <= num; i +=1 {}

	return (i -1)*(i -1) == num
}