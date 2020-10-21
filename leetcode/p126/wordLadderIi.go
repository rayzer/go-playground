package main

import "fmt"

/*
 * @lc app=leetcode id=126 lang=golang
 *
 * [126] Word Ladder II
 */

// @lc code=start

type step struct {
	current     string
	step        int
	wordListMap map[string]bool
	history     []string
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var results [][]string
	dict := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	wordListMap := make(map[string]bool)
	for _, s := range wordList {
		wordListMap[s] = true
	}

	queue := []step{{beginWord, 1, wordListMap, []string{beginWord}}}
	for len(queue) != 0 {
		fmt.Println(queue)

		t := queue[0]
		queue = queue[1:]

		if t.current == endWord {
			fmt.Println("\n\n\nadd result............", t)
			results = append(results, t.history)
			continue
		}

		for i := 0; i < len(t.current); i++ {
			for _, o := range dict {
				temp := t.current[:i] + o + t.current[i+1:]
				if _, ok := t.wordListMap[temp]; ok {
					newMap := make(map[string]bool)
					for k, v := range t.wordListMap {
						newMap[k] = v
					}
					delete(newMap, temp)

					queue = append(queue, step{temp, t.step + 1, newMap, append(t.history, temp)})
				}
			}
		}
	}
	return results
}

// @lc code=end

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
