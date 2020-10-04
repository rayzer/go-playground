package main

import (
	"fmt"
)

type ByteSlice []byte

func (s ByteSlice) String() string {
	return fmt.Sprintf("%v", []byte(s))
}

/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start

type step struct {
	current string
	step    int
}

func minMutation(start string, end string, bank []string) int {
	bankMap := make(map[string]bool)
	for _, val := range bank {
		bankMap[val] = true
	}
	//options := []string{"A", "C", "G", "T"}
	optionsMap := make(map[string][]string)
	optionsMap["A"] = []string{"C", "G", "T"}
	optionsMap["C"] = []string{"A", "G", "T"}
	optionsMap["G"] = []string{"A", "C", "T"}
	optionsMap["T"] = []string{"A", "C", "G"}

	var queue []step
	queue = append(queue, step{start, 0})
	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]
		if t.current == end {
			return t.step
		}
		for i := 0; i < len(t.current); i++ {
			for _, o := range optionsMap[string(t.current[i])] {
				temp := t.current[:i] + o + t.current[i+1:]
				if _, ok := bankMap[temp]; ok {
					delete(bankMap, temp)
					queue = append(queue, step{temp, t.step + 1})
				}
			}
		}
	}
	return -1
}

// @lc code=end

func main() {
	start := "AAAAACCC"
	end := "AACCCCCC"
	bank := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC", "AAAAACCC"}
	fmt.Println(minMutation(start, end, bank))
}
