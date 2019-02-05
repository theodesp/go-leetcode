package trie

import "testing"

func TestWordDictionary_Search(t *testing.T)  {
	trie := DictConstructor()

	trie.AddWord("bad")
	trie.AddWord("dad")
	trie.AddWord("mad")

	if trie.Search("pad") {
		t.Fail()
	}

	if !trie.Search("bad") {
		t.Fail()
	}

	if !trie.Search(".ad") {
		t.Fail()
	}

	if trie.Search("s..") {
		t.Fail()
	}

	if !trie.Search("b.d") {
		t.Fail()
	}

	if trie.Search("b...") {
		t.Fail()
	}


}