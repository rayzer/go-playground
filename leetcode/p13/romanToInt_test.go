package p13

import "testing"

func Test_romanToInt(t *testing.T) {
	if 3 != romanToInt("III") {
		t.Error()
	}
	if 4 != romanToInt("IV") {
		t.Error()
	}
	if 9 != romanToInt("IX") {
		t.Error()
	}
	if 58 != romanToInt("LVIII") {
		t.Error()
	}
	if 1994 != romanToInt("MCMXCIV") {
		t.Error(romanToInt("MCMXCIV"))
	}
	if 621 != romanToInt("DCXXI") {
		t.Error(romanToInt("DCXXI"))
	}
}
