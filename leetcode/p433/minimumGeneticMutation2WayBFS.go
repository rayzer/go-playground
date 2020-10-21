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

func minMutation2WayBFS(start string, end string, bank []string) int {
	bankMap := make(map[string]bool)
	for _, val := range bank {
		bankMap[val] = true
	}

	if _, ok := bankMap[end]; !ok {
		return -1
	}

	//options := []string{"A", "C", "G", "T"}
	optionsMap := make(map[string][]string)
	optionsMap["A"] = []string{"C", "G", "T"}
	optionsMap["C"] = []string{"A", "G", "T"}
	optionsMap["G"] = []string{"A", "C", "T"}
	optionsMap["T"] = []string{"A", "C", "G"}

	mutate := 0
	fromBegin := map[string]struct{}{start: {}}
	fromEnd := map[string]struct{}{end: {}}

	for len(fromBegin) > 0 && len(fromEnd) > 0 {
		mutate++
		temp := make(map[string]struct{})
		for k := range fromBegin {
			for i := 0; i < len(k); i++ {
				options := optionsMap[string(k[i])]
				for j := 0; j < len(options); j++ {
					if string(k[i]) != options[j] {
						tmp := k[:i] + options[j] + k[i+1:]

						if _, ok := fromEnd[tmp]; ok {
							return mutate
						}

						if _, ok := bankMap[tmp]; ok {
							temp[tmp] = struct{}{}
							delete(bankMap, tmp)
						}
					}
				}
			}
		}
		fromBegin = temp
		if len(fromBegin) > len(fromEnd) {
			fromBegin, fromEnd = fromBegin, fromEnd
		}
	}
	return -1
}

// @lc code=end

func main() {
	start := "AAAAACCC"
	end := "AACCCCCC"
	bank := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC", "AAAAACCC"}
	fmt.Println(minMutation2WayBFS(start, end, bank))
	fmt.Println(minMutation2WayBFS("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
	fmt.Println(minMutation2WayBFS("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))
}
