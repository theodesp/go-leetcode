package uncategorized

/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
 */

 /*
 Solution:
 Get lengths of a and b (l1 and l2)
 Init curry variable
 Start from the last element of a and b
 While we haven't reached the start of a and b
 Check if a(l1)
 If curry == 1 then prepend 1 in result

 Time Complexity: O(max(len(a),len(b)) = O(n)
 Space Complexity: O(max(len(a),len(b)) = O(n)
  */
func AddBinary(a string, b string) string {
	res := ""
	curry := 0
	l1 := len(a) - 1
	l2 := len(b) - 1
	for l1 >= 0 || l2 >= 0 {
		var left = 0
		var right = 0
		if l1 >= 0 {
			if rune(a[l1]) == '1' {
				left = 1
			}
			l1 -= 1
		}
		if l2 >= 0 {
			if rune(b[l2]) == '1' {
				right = 1
			}
			l2 -= 1
		}
		added := curry + left + right
		switch added {
		case 0:
			res = "0" + res
			curry = 0
		case 1:
			res = "1" + res
			curry = 0
		case 2:
			res = "0" + res
			curry = 1
		case 3:
			res = "1" + res
			curry = 1
		}
	}

	if curry == 1 {
		res = "1" + res
	}
	return res
}
