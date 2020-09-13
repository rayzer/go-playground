package main

import (
	"os"
	"strings"
)

func maxScore(s string) int {
	if len(s) > 500 && len(s) < 2 {
		os.Exit(-1)
	}
	var maxScore int
	for i := 1; i < len(s); i++ {
		score := strings.Count(s[0:i], "0") + strings.Count(s[i:], "1")
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}
