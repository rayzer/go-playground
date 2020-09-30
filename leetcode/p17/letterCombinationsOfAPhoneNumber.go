package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {

	letters := make(map[int]string)
	letters[2] = "abc"
	letters[3] = "edf"
	letters[4] = "ghi"
	letters[5] = "jkl"
	letters[6] = "mno"
	letters[7] = "pqrs"
	letters[8] = "tuv"
	letters[9] = "wxyz"
	var result []string
	if len(digits) == 0 {
		return []string{}
	}
	_letterCombinations(&result, digits, "", 0, &letters)
	return result
}

func _letterCombinations(result *[]string, digits string, s string, idx int, letters *map[int]string) {
	if idx == len(digits) {
		*result = append(*result, s)
		return
	}
	//handle
	pos, _ := strconv.Atoi(string(digits[idx]))
	options := (*letters)[pos]
	for i := 0; i < len(options); i++ {
		_letterCombinations(result, digits, s+string(options[i]), idx+1, letters)
	}
}

// @lc code=end
func main() {
	fmt.Println(letterCombinations("234"))
}
