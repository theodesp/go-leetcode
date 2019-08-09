package uncategorized

import "testing"

func TestFirstUniqChar(t * testing.T)  {
	if firstUniqChar("leetcode") != 0 {
		t.Fail()
	}

	if firstUniqChar("loveleetcode") != 2 {
		t.Fail()
	}
}