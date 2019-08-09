package uncategorized

import "testing"

func TestLongestSuffixPrefix(t * testing.T)  {
	if LongestSuffixPrefix("aabcdaabc") != 4 {
		t.Fail()
	}
}