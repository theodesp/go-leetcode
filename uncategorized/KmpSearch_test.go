package uncategorized

import "testing"

func TestKmpSearch(t * testing.T)  {
	if KmpSearch("ll", "hello") != 2 {
		t.Fail()
	}

	if KmpSearch("ABABCABAB", "ABABDABACDABABCABAB") != 10 {
		t.Fail()
	}
}