package p283

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

//[0,1,0,3,12]
//[1,3,12,0, 0]

// @lc code=start
func moveZeroes(nums []int) {
	j := 0
	i := 1
	for i < len(nums) {
		if nums[j] == 0 {
			if nums[i] == 0 {
				i++
			} else { //found non 0
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}
}

//func moveZeroes(nums []int) {
//	j := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] != 0 {
//			if i != j {
//				nums[i], nums[j] = nums[j], nums[i]
//				j++
//			} else {
//				j++
//			}
//		}
//	}
//}

// @lc code=end
