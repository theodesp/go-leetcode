package trie

import "testing"

func TestTrie_Search(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Fail()
	}

	if trie.Search("app") {
		t.Fail()
	}

	trie.Insert("app")
	if !trie.Search("app") {
		t.Fail()
	}
}

func TestTrie_StartsWith(t *testing.T) {
	trie := Constructor()
	trie.Insert("cattle")

	if !trie.StartsWith("cat") {
		t.Fail()
	}

	if trie.StartsWith("cap") {
		t.Fail()
	}
}