package uncategorized

import "math"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	lookup := make(map[uint8]int)
	for i, j := 0, 0; j < n; j +=1 {
		if val, ok := lookup[s[j]]; ok { // found
			i = int(math.Max(float64(val), float64(i)))
		}
		ans = int(math.Max(float64(ans), float64(j - i + 1)))
		lookup[s[j]] = j + 1
	}

	return ans
}
