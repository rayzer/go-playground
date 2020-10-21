package main

import "fmt"

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

//不能挨着取
//dp subproblem
//dp array a[i] : 能偷到的最大值
// a[i][0,1] : 0: 偷 i；1：不偷 i
//dp function
// 不偷 i 的最大值：a[i][0] = Max(a[i-1][0], a[i-1][1])
// 偷 i 的最大值： a[i][1] = a[i-1][0] + nums[i]
// 结果 max(a[n][0], a[n][1])

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	dp := [2][]int{}
	dp[0] = append(dp[0], 0)
	dp[1] = append(dp[1], nums[0])
	fmt.Println(dp)

	for i := 1; i < len(nums); i++ {
		dp[0] = append(dp[0], max(dp[0][i-1], dp[1][i-1]))
		dp[1] = append(dp[1], dp[0][i-1]+nums[i])
	}
	return max(dp[0][n-1], dp[1][n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//a[i]: 0..i 天， nums[i]必偷的最大值
//a[i] = max(a[i-1], a[i-2] + nums[i])
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := []int{nums[0]}
	dp = append(dp, max(nums[0], nums[1]))
	result := max(dp[0], dp[1])

	for i := 2; i < len(nums); i++ {
		dp = append(dp, max(dp[i-1], dp[i-2]+nums[i]))
		result = max(result, dp[i])
	}
	return result
}

// @lc code=end
func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))     //4
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))  //12
	fmt.Println(rob2([]int{1, 2, 3, 1}))    //4
	fmt.Println(rob2([]int{2, 7, 9, 3, 1})) //12
}
