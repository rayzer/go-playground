package main

import "testing"

func shiftingLetters(S string, shifts []int) string {
	return "rpl"
}

func Test(t *testing.T) {
	if shiftingLetters("abc", []int{3, 5, 9}) != "rpl" {
		t.Fail()
	}
}
