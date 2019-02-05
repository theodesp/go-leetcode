package trie

/*
Design a data structure that supports the following two operations:

void addWord(word)
bool search(word)
search(word) can search a literal word or a regular expression string containing
only letters a-z or .. A . means it can represent any one letter.

Example:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
Note:
You may assume that all words are consist of lowercase letters a-z.
 */

type WordDictionary struct {
	children map[rune]*WordDictionary
	isTerminalWord bool
}

/** Initialize your data structure here. */
func DictConstructor() WordDictionary {
	return WordDictionary{
		children: make(map[rune]*WordDictionary),
		isTerminalWord: false,
	}
}

/** Adds a word into the data structure. */
func (t *WordDictionary) AddWord(word string) {
	t.Insert(word)
}

/** Inserts a word into the trie. */
func (t *WordDictionary) Insert(word string) {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[rune(ch)]; !ok {
			trie := DictConstructor()
			curr.children[rune(ch)] = &trie
		}
		curr = curr.children[rune(ch)]
	}

	curr.isTerminalWord = true
}

/** Returns if the word is in the data structure.
A word could contain the dot character '.' to represent any one letter. */
func (t *WordDictionary) Search(word string) bool {
	curr := t
	for i, ch := range word {
		if rune(ch) == '.' {
			isMatched := false
			for _, v := range curr.children {
				if v.Search(word[i+1:]) {
					isMatched = true
				}
			}
			return isMatched
		} else if _, ok := curr.children[rune(ch)]; !ok {
			return false
		}
		curr = curr.children[rune(ch)]
	}

	return len(curr.children) == 0 || curr.isTerminalWord
}
