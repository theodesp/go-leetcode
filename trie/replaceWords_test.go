package trie

import "testing"

func Test_ReplaceWords(t *testing.T) {
	result := ReplaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery")
	if result != "the cat was rat by the bat" {
		t.Fail()
	}
}
