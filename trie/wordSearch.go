package trie

/**
Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

Example:

Input:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]

Output: ["eat","oath"]
Note:
You may assume that all inputs are consist of lowercase letters a-z.
 */


const mark = byte('#')

func dfs(board [][]byte, i int, j int, trie Trie, result *[]string, curr string)  {
	c := board[i][j]
	_, ok := trie.Children[rune(c)]
	if board[i][j] == mark || !ok  {
		return
	}
	trie = *trie.Children[rune(c)]
	if trie.isTerminalWord {
		*result = append(*result, curr + string(c))
	}
	board[i][j] = mark
	if i > 0 {
		dfs(board, i-1, j, trie, result, curr + string(c))
	}
	if j > 0 {
		dfs(board, i, j-1, trie, result, curr + string(c))
	}
	if i < len(board) - 1 {
		dfs(board, i+1, j, trie, result, curr + string(c))
	}
	if j < len(board[0]) - 1 {
		dfs(board, i, j+1, trie, result, curr + string(c))
	}
	board[i][j] = c
}

func removeDuplicates(array []string) (result []string) {
	seen := map[string]bool{}

	for _, value := range array {
		if seen[value] == false {
			seen[value] = true
			result = append(result, value)
		}
	}

	return result
}

func FindWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	result := make([]string, 0)
	for i, row := range board {
		for j := range row {
			dfs(board, i, j, trie, &result, "")
		}
	}

	return removeDuplicates(result)
}