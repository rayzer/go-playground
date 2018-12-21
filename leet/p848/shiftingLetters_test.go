//https://leetcode.com/contest/weekly-contest-88/problems/shifting-letter
//https://golang.org/ref/spec#Conversions
package main

import (
	"testing"
)

func shiftingLetters(S string, shifts []int) string {
	for i := len(shifts) - 2; i >= 0; i-- {
		shifts[i] = shifts[i] + shifts[i+1]
	}
	result := []byte(S)
	for i := 0; i < len(result); i++ {
		result[i] = byte((int(result[i])-'a'+shifts[i])%26 + 'a')
	}
	return string(result)
}

func Test(t *testing.T) {
	if shiftingLetters("abc", []int{3, 5, 9}) != "rpl" {
		t.Fail()
	}
	if shiftingLetters("bad", []int{10, 20, 30}) != "jyh" {
		t.Fail()
	}
}
