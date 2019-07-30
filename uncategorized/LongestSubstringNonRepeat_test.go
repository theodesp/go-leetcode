package uncategorized

import "testing"

func TestLongestSubstringNonRepeat(t *testing.T) {
	//if lengthOfLongestSubstring("") != 0 {
	//	t.Fail()
	//}
	//if lengthOfLongestSubstring("a") != 1 {
	//	t.Fail()
	//}
	//if lengthOfLongestSubstring("vvvv") != 1 {
	//	t.Fail()
	//}
	//if lengthOfLongestSubstring("dvdf") != 3 {
	//	t.Fail()
	//}
	//if lengthOfLongestSubstring("abcdefg") != 7 {
	//	t.Fail()
	//}
	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Errorf("Failed when input=%s", "pwwkew")
	}
}