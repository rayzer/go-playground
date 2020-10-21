package main

import "fmt"

/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */

/*
1. word遍历 --》 board search
2. trie
	words 构建 trie
	board DFS
*/

// @lc code=start
type Trie struct {
	links [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for _, c := range word {
		idx := c - 'a'
		if curr.links[idx] == nil {
			curr.links[idx] = &Trie{}
		}
		curr = curr.links[idx]
	}
	curr.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this
	for _, c := range word {
		idx := c - 'a'
		if curr.links[idx] == nil {
			return false
		}
		curr = curr.links[idx]
	}
	return curr.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, c := range prefix {
		idx := c - 'a'
		if curr.links[idx] == nil {
			return false
		}
		curr = curr.links[idx]
	}
	return true
}

func findWords(board [][]byte, words []string) []string {
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	resultSet := make(map[string]bool)

	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	m, n := len(board), len(board[0])

	var dfs func(board [][]byte, i, j int, c string, trie *Trie)
	dfs = func(board [][]byte, i, j int, c string, trie *Trie) {
		c += string(board[i][j])
		if trie.Search(c) {
			resultSet[c] = true
		}
		var tmp byte
		tmp, board[i][j] = board[i][j], '#'
		if trie.StartsWith(c) {
			for k := 0; k < 4; k++ {
				x, y := i+dx[k], j+dy[k]
				if (x >= 0 && x < m) && (y >= 0 && y < n) && //x, y 在 board 上
					board[x][y] != '#' && trie.StartsWith(c+string(board[x][y])) {
					dfs(board, x, y, c, trie)
				}
			}
		}
		board[i][j] = tmp
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if trie.StartsWith(string(board[i][j])) {
				dfs(board, i, j, "", &trie)
			}
		}
	}

	var result []string
	for s := range resultSet {
		result = append(result, s)
	}
	return result
}

// @lc code=end

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words)) //"eat", "oath"
	fmt.Println(findWords([][]byte{{'a'}, {'a'}}, []string{"a"}))
}
