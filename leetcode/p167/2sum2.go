package p167

//https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := len(numbers) - 1; j > i; j-- {
			if numbers[j]+numbers[i] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}
