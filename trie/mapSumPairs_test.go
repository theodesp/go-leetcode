package trie

import (
	"testing"
)

func TestMapSum_Sum(t *testing.T) {
	trie := MapSumConstructor()
	trie.Insert("apple", 3)

	if trie.Sum("ap") != 3 {
		t.Fail()
	}

	trie.Insert("app", 2)

	if trie.Sum("ap") != 5 {
		t.Fail()
	}
}
