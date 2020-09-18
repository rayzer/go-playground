package p07

import (
	"math"
	"strconv"
)

//0 ms, faster than 100.00%
func reverse(x int) int {
	var stringNoSign string
	if x < 0 {
		stringNoSign = strconv.Itoa(x)[1:]
	} else {
		stringNoSign = strconv.Itoa(x)
	}

	var i = len(stringNoSign) - 1
	var reverseString string
	for i > -1 {
		reverseString += string(stringNoSign[i])
		i--
	}
	reversed, _ := strconv.ParseInt(reverseString, 10, 64)

	if x < 0 {
		reversed = -reversed
	}
	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}

	return int(reversed)
}

//Runtime: 4 ms, faster than 43.46%
/*func reverse(x int) int {
	t := x
	if x < 0 {
		t = -1 * x
	}

	var reversed int
	for t != 0 {
		lastDigit := t % 10
		t = t / 10
		reversed = reversed*10 + lastDigit
	}

	if reversed > math.MaxInt32 {
		return 0
	}

	if x < 0 {
		reversed = -1 * reversed
	}
	return reversed
}*/

//Runtime: 4 ms, faster than 43.46%
/*func reverse(x int) int {

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	var reversed int
	reversed = 0
	for x != 0 {
		lastDigit := x % 10
		x = x / 10
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && lastDigit > 7) {
			return 0
		}
		if reversed < math.MinInt32/10 || (reversed == math.MinInt32/10 && lastDigit < -8) {
			return 0
		}
		reversed = reversed*10 + lastDigit
	}

	return reversed
}*/

func main() {
	println(reverse(
		1534236469))
}
