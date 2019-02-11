package trie

type Trie struct {
	children       map[rune]*Trie
	isTerminalWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		children:       make(map[rune]*Trie),
		isTerminalWord: false,
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			trie := Constructor()
			curr.children[ch] = &trie
		}
		curr = curr.children[ch]
	}
	curr.isTerminalWord = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			return false
		}
		curr = curr.children[ch]
	}

	return len(curr.children) == 0 || curr.isTerminalWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	curr := t
	for _, ch := range prefix {
		if _, ok := curr.children[ch]; !ok {
			return false
		}
		curr = curr.children[ch]
	}

	return true
}
