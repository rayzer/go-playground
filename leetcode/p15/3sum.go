package p15

import "sort"

//https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	var resultSet [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				resultSet = append(resultSet, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return resultSet
}

//func threeSum(nums []int) [][]int {
//	var resultSet [][]int
//	sort.Ints(nums)
//	for i, v := range nums {
//		if i > 0 && v == nums[i-1] {
//			continue
//		}
//
//		l, r := i+1, len(nums)-1
//		for l < r {
//			sum := v + nums[l] + nums[r]
//			if sum > 0 {
//				r--
//			} else if sum < 0 {
//				l++
//			} else {
//				resultSet = append(resultSet, []int{v, nums[l], nums[r]})
//				l++
//				//skip duplicate numbers
//				for nums[l] == nums[l-1] && l < r {
//					l++
//				}
//			}
//		}
//	}
//	return resultSet
//}
