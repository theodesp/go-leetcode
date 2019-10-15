package uncategorized

/*
Given a string, Check if characters of the given string can be
rearranged to form a palindrome.
 */

/* Solution: We know that a palidrome has at most 1 odd number of char occurences.
 We construct a char count array and check if more that one odd char count is found.
 */
func CanFormPalindrome(s string) bool {
	charCount := make(map[rune]int)
	for _, elem := range s {
		charCount[elem] += 1
	}
	oddsCount := 0
	for _, val := range charCount {
		if val % 2 == 1 {
			oddsCount += 1
			if oddsCount > 1 {
				return false
			}
		}
	}
	return true
}