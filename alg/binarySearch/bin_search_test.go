package binarySearch

import "testing"

func Test(t *testing.T) {
	data := []int{1, 2, 4, 5, 7, 8, 9}
	if 6 != binsearch(data, len(data), 9) {
		t.Error("position shall be 6, but: ")
	}

	if 1 != binsearch(data, len(data), 2) {
		t.Error("position shall be 1, but: ")
	}
}
