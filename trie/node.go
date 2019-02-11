package trie

const Cases = 26

// Waste of space
type ArrayTrieNode struct {
	children *[Cases]ArrayTrieNode
}

func NewArrayTrieNode() *ArrayTrieNode {
	return &ArrayTrieNode{
		children: &[Cases]ArrayTrieNode{},
	}
}

// More flexible
type MapTrieNode struct {
	children map[rune]MapTrieNode
}

func NewMapTrieNode() *MapTrieNode {
	return &MapTrieNode{
		children: make(map[rune]MapTrieNode),
	}
}
