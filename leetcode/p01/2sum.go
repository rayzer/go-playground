package p01

//https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		completionNumber := target - nums[i]
		if _, ok := m[completionNumber]; ok {
			return []int{m[completionNumber], i}
		}
		m[nums[i]] = i
	}
	return nil
}
