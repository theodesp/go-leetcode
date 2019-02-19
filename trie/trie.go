package trie

type Trie struct {
	Children       map[rune]*Trie
	isTerminalWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Children:       make(map[rune]*Trie),
		isTerminalWord: false,
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	curr := t
	for _, ch := range word {
		if _, ok := curr.Children[ch]; !ok {
			trie := Constructor()
			curr.Children[ch] = &trie
		}
		curr = curr.Children[ch]
	}
	curr.isTerminalWord = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	curr := t
	for _, ch := range word {
		if _, ok := curr.Children[ch]; !ok {
			return false
		}
		curr = curr.Children[ch]
	}

	return len(curr.Children) == 0 || curr.isTerminalWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	curr := t
	for _, ch := range prefix {
		if _, ok := curr.Children[ch]; !ok {
			return false
		}
		curr = curr.Children[ch]
	}

	return true
}
