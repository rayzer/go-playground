package main

import (
	"fmt"
	"reflect"
)

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
//1: using hashtable
//2:sort

//1. 16 ms, faster than 19.32%

/*func isAnagram(s string, t string) bool {
	ss := []byte(s)
	st := []byte(t)
	sort.Slice(ss, func(i int, j int) bool { return ss[i] < ss[j] })
	sort.Slice(st, func(i int, j int) bool { return st[i] < st[j] })
	fmt.Println(ss)
	fmt.Println(st)
	return reflect.DeepEqual(ss, st)
}*/

//8 ms, faster than 54.03%
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 && len(t) == 0 {
		return true
	}
	ss := []byte(s)
	st := []byte(t)
	ss_map := make(map[byte]int)
	st_map := make(map[byte]int)
	for i := 0; i < len(ss); i++ {
		ss_map[ss[i]] += 1
	}
	for i := 0; i < len(st); i++ {
		st_map[st[i]] += 1
	}
	return reflect.DeepEqual(ss_map, st_map)
}

// @lc code=end

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
