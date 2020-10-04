package main

import "fmt"

/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start

type step struct {
	current string
	step    int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	wordListMap := make(map[string]bool)
	for _, s := range wordList {
		wordListMap[s] = true
	}

	queue := []step{{beginWord, 1}}
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]

		if t.current == endWord {
			return t.step
		}

		for i := 0; i < len(t.current); i++ {
			for _, o := range dict {
				temp := t.current[:i] + o + t.current[i+1:]
				if _, ok := wordListMap[temp]; ok {
					delete(wordListMap, temp)
					queue = append(queue, step{temp, t.step + 1})
				}
			}
		}
	}

	return 0
}

// @lc code=end
func main() {
	//hit -> hot -> dot -> dog -> cog
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
