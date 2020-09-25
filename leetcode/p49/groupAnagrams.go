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
	dict := make(map[string][]string)
	sortString := func(s string) string {
		tmp := strings.Split(s, "")
		sort.Strings(tmp)
		return strings.Join(tmp, "")
	}
	for _, val := range strs {
		ss := sortString(val)
		if _, ok := dict[ss]; ok {
			dict[ss] = append(dict[ss], val)
		} else {
			dict[ss] = []string{val}
		}
	}

	var result [][]string
	for _, val := range dict {
		result = append(result, val)
	}
	return result
}

// @lc code=end

//1：将每个 string 排序，再插入 map
//2：将每个 string 处理成统计值，在以统计值 map

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
