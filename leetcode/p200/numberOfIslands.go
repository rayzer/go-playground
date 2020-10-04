package main

import "fmt"

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			for k := 0; k < len(grid); k++ {
				fmt.Printf("%c\n", grid[k])
			}
			fmt.Println()
			fmt.Println()
			islands += sinkIsland(i, j, grid)
		}
	}
	return islands
}

func sinkIsland(i int, j int, grid [][]byte) int {
	if grid[i][j] == '0' {
		return 0
	}

	fmt.Println("sink ij : ", i, j)

	grid[i][j] = '0'
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1} //( -1, 0) 左边点 , (1,0) 右边点, (0, -1) 上方点, （0，1）下方点

	//检查上下左右是否为水0，否则 sink
	for k := 0; k < len(dx); k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
			if grid[x][y] == '0' {
				continue
			}
			fmt.Println("sink xy : ", x, y)
			sinkIsland(x, y, grid)
		}
	}
	return 1
}

// @lc code=end

func main() {
	grid := [][]byte{{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
}
