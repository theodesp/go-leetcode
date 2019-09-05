package uncategorized

import (
	"math"
)

/*
Given a non-empty string s, you may delete at most one character.
Judge whether you can make it a palindrome.

Example 1:
Input: "aba"
Output: True
Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
Note:
The string will only contain lowercase characters a-z.
The maximum length of the string is 50000.
 */

 /*
 Solution
  */

func ValidPalindrome(s string) bool {
	l := len(s)
	m := int(math.Floor(float64(len(s) / 2)))
	for i:= 0;i <= m; i+= 1 {
		if s[i] != s[l-i-1] {
			return IsPalindrome(s[:i] + s[i+1:]) || IsPalindrome(s[:i+1] + s[l-i:])
		}
	}
	return true
}

func IsPalindrome(s string) bool  {
	l := len(s)
	m := int(math.Floor(float64(len(s) / 2)))
	if len(s) == 1 {
		return false
	}
	if len(s) == 2 && s[0] != s[1] {
		return false
	}
	for i:= 0;i< m; i+= 1 {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
