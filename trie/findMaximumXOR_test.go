package trie

import "testing"

func TestFindMaximumXOR(t *testing.T) {
	if FindMaximumXOR([]int{3, 10, 5, 25, 2, 8}) != 28 {
		t.Fail()
	}
}
