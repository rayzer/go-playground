package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
// @lc code=start
func groupAnagrams(strs []string) [][]string {
	patternMap := map[[26]byte]int{}
	res := [][]string{}

	for _, str := range strs {
		pattern := [26]byte{}
		for _, l := range str {
			pattern[l-'a']++
		}

		if idx, found := patternMap[pattern]; found {
			res[idx] = append(res[idx], str)
		} else {
			res = append(res, []string{str})
			patternMap[pattern] = len(res) - 1
		}
	}
	return res
}

// @lc code=end

//1：将每个 string 排序，再插入 map
//2：将每个 string 处理成统计值，在以统计值 map

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
