package p167

//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {
	//for i := 0; i < len(numbers); i++ {
	//	for j := len(numbers) - 1; j > i; j-- {
	//		if numbers[j]+numbers[i] == target {
	//			return []int{i + 1, j + 1}
	//		}
	//	}
	//}

	//[2,7,11,15]
	//9
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}
