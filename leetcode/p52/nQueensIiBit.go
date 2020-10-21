package main

import "fmt"

/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start
func totalNQueensB(n int) int {
	var count int
	size := (1 << n) - 1 //可以填皇后的位置，把所有有效位置1: 如8个1

	var solve func(row, ld, rd int)
	solve = func(row, ld, rd int) {
		//递归终止，当前行已无位置可放（全1）
		if row == size {
			count++
			return
		}
		pos := size & (^(row | ld | rd)) //当前层能放皇后的所有位置（所有被占位置取反与 size 求与，得到0

		//当前层还有能放的位置（不全为0），分别放上皇后下探
		for pos != 0 {
			p := pos & (-pos)                  //取到最低位的1
			pos -= p                           //让最低位的1为0，放上皇后
			solve(row|p, (ld|p)<<1, (rd|p)>>1) //下探，该位置的下一行同一位置（列），该位置的下一行左斜、右斜方向，分别向左右展开占位（撇捺）
		}
	}

	solve(0, 0, 0)
	return count
}

// @lc code=end
func main() {
	fmt.Println(totalNQueensB(4))
}
