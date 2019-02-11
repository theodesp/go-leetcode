package trie

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

/*
Given a non-empty array of numbers, a0, a1, a2, … , an-1,
where 0 ≤ ai < 231.

Find the maximum result of ai XOR aj, where 0 ≤ i, j < n.

Could you do this in O(n) runtime?

Example:

Input: [3, 10, 5, 25, 2, 8]

Output: 28

Explanation: The maximum result is 5 ^ 25 = 28.
*/

type BinaryTrie struct {
	children       map[rune]*BinaryTrie
	isTerminalWord bool
}

/** Initialize your data structure here. */
func BConstructor() BinaryTrie {
	return BinaryTrie{
		children:       make(map[rune]*BinaryTrie),
		isTerminalWord: false,
	}
}

/** Inserts a word into the trie. */
func (t *BinaryTrie) Insert(word string) {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			trie := BConstructor()
			curr.children[ch] = &trie
		}
		curr = curr.children[ch]
	}
	curr.isTerminalWord = true
}

/** Returns if the word is in the trie. */
func (t *BinaryTrie) Search(word string) bool {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			return false
		}
		curr = curr.children[ch]
	}

	return len(curr.children) == 0 || curr.isTerminalWord
}

func max(nums []int) int {
	max := nums[0]
	for _, v := range nums[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

func lpad(s string, pad string, plength int) string {
	for i := len(s); i < plength; i++ {
		s = pad + s
	}
	return s
}

func FindMaximumXOR(nums []int) int {
	offset := int(math.Log2(float64(max(nums)))) + 1

	t := BConstructor()
	for _, num := range nums {
		num := lpad(strconv.FormatInt(int64(num), 2), "0", offset)
		t.Insert(num)
	}

	curMax := 0
	for _, num := range nums {
		acc := bytes.NewBufferString("")
		num := lpad(strconv.FormatInt(int64(num), 2), "0", offset)

		curr := t
		for _, ch := range num {
			if ch == '0' {
				if _, ok := curr.children['1']; ok {
					acc.WriteString("1")
					curr = *curr.children['1']
					continue
				}
				if _, ok := curr.children['0']; ok {
					acc.WriteString("0")
					curr = *curr.children['0']
					continue
				}
			}
			if ch == '1' {
				if _, ok := curr.children['0']; ok {
					acc.WriteString("1")
					curr = *curr.children['0']
					continue
				}
				if _, ok := curr.children['1']; ok {
					acc.WriteString("0")
					curr = *curr.children['1']
					continue
				}
			}
		}

		s := acc.String()
		i, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		if i > int64(curMax) {
			curMax = int(i)
		}
	}

	return curMax
}
