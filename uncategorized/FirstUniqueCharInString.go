package uncategorized

/*
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
 */

 // Solution
 // 2 iterations.
 // First mark each char count and second will find the first one that has 1 occurence
 // Time complexity O(n)
 // Space complexity O(n)
func firstUniqChar(s string) int {
	lookup := make(map[rune]int)
	for _, char := range s {
		if val, ok := lookup[char]; ok {
			lookup[char] = val + 1
		} else {
			lookup[char] = 1
		}
	}
	for pos, char := range s {
		if lookup[char] == 1 {
			return pos
		}
	}

	return -1
}