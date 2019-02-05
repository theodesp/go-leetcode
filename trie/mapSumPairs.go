package trie

/*
Implement a MapSum class with insert, and sum methods.

For the method insert, you'll be given a pair of (string, integer).
The string represents the key and the integer represents the value.
If the key already existed, then the original key-value pair will be overridden to
the new one.

For the method sum, you'll be given a string representing the prefix,
and you need to return the sum of all the pairs' value whose key starts with the prefix.

Example 1:
Input: insert("apple", 3), Output: Null
Input: sum("ap"), Output: 3
Input: insert("app", 2), Output: Null
Input: sum("ap"), Output: 5
 */

type MapSum struct {
	children       map[rune]*MapSum
	value          int
	isTerminalWord bool
}

/** Initialize your data structure here. */
func MapSumConstructor() MapSum {
	return MapSum{
		children:       make(map[rune]*MapSum),
		value:          0,
		isTerminalWord: false,
	}
}

/** Inserts a word into the trie. */
func (t *MapSum) Insert(word string, value int) {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[rune(ch)]; !ok {
			trie := MapSumConstructor()
			curr.children[rune(ch)] = &trie
		}
		curr = curr.children[rune(ch)]
	}
	curr.isTerminalWord = true
	curr.value = value
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *MapSum) StartsWith(prefix string) bool {
	curr := t
	for _, ch := range prefix {
		if _, ok := curr.children[rune(ch)]; !ok {
			return false
		}
		curr = curr.children[rune(ch)]
	}

	return true
}

func (t *MapSum) Sum(prefix string) int {
	curr := t
	for _, ch := range prefix {
		if _, ok := curr.children[rune(ch)]; !ok {
			return 0
		}
		curr = curr.children[rune(ch)]
	}

	return SumChildren(*curr)
}

func SumChildren(t MapSum) int {
	curr := t
	sum := t.value
	for k := range curr.children {
		sum = sum + SumChildren(*curr.children[k])
	}

	return sum
}
