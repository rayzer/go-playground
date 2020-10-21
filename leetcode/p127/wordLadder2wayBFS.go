package main

import "fmt"

/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

// @lc code=start

func ladderLength2WayBFS(beginWord string, endWord string, wordList []string) int {
	dict := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	wordListMap := make(map[string]struct{})
	for _, s := range wordList {
		wordListMap[s] = struct{}{}
	}

	if _, ok := wordListMap[endWord]; !ok {
		return 0
	}

	step := 1
	fromBegin := map[string]struct{}{beginWord: {}}
	fromEnd := map[string]struct{}{endWord: {}}

	for len(fromBegin) > 0 && len(fromEnd) > 0 {
		step++
		tempSet := make(map[string]struct{})
		for k := range fromBegin {
			for i := 0; i < len(k); i++ {
				for _, d := range dict {
					if d != string(k[i]) {
						tmp := k[:i] + d + k[i+1:]
						if _, ok := fromEnd[tmp]; ok {
							return step
						}
						if _, ok := wordListMap[tmp]; ok {
							tempSet[tmp] = struct{}{}
							delete(wordListMap, tmp)
						}
					}
				}
			}
		}
		fromBegin = tempSet
		if len(fromBegin) > len(fromEnd) {
			fromBegin, fromEnd = fromEnd, fromBegin
		}
	}
	return 0
}

// @lc code=end
func main() {
	//hit -> hot -> dot -> dog -> cog
	fmt.Println(ladderLength2WayBFS("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	//fmt.Println(ladderLength2WayBFS("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	//fmt.Println(ladderLength2WayBFS("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}
