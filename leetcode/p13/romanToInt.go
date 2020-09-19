package p13

/*
I 1
V 5
X 10
L 50
C 100
D 500
M 1000
III 3
IV 4
*/
//8 ms, faster than 67.90%
func romanToInt(s string) int {
	result := 0
	var pre byte
	for _, v := range s {
		cur := byte(v)
		switch cur {
		case 'I':
			result += 1 + checkNextByte(cur, pre)
		case 'V':
			result += 5 + checkNextByte(cur, pre)
		case 'X':
			result += 10 + checkNextByte(cur, pre)
		case 'L':
			result += 50 + checkNextByte(cur, pre)
		case 'C':
			result += 100 + checkNextByte(cur, pre)
		case 'D':
			result += 500 + checkNextByte(cur, pre)
		case 'M':
			result += 1000 + checkNextByte(cur, pre)
		}
		pre = cur
	}
	return result
}

func checkNextByte(cur byte, pre byte) int {
	switch {
	case (cur == 'V' || cur == 'X') && pre == 'I':
		return -2
	case (cur == 'L' || cur == 'C') && pre == 'X':
		return -20
	case (cur == 'D' || cur == 'M') && pre == 'C':
		return -200
	}
	return 0
}

//12 ms, faster than 37.93%
/*func romanToInt(s string) int {
	input := []byte(s)
	result := 0

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case 'I':
			if i+1 < len(input) && (input[i+1] == 'V' || input[i+1] == 'X') {
				result--
			} else {
				result++
			}
		case 'V':
			result += 5
		case 'X':
			if i+1 < len(input) && (input[i+1] == 'L' || input[i+1] == 'C') {
				result -= 10
			} else {
				result += 10
			}
		case 'L':
			result += 50
		case 'C':
			if i+1 < len(input) && (input[i+1] == 'D' || input[i+1] == 'M') {
				result -= 100
			} else {
				result += 100
			}
		case 'D':
			result += 500
		case 'M':
			result += 1000
		}
	}
	return result
}*/
