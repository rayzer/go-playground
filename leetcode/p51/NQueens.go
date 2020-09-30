package main

import "fmt"

func solveNQueens(n int) [][]string {
	var queens [][2]int
	var result [][]string
	var solutions [][][2]int //2: 坐标
	place(n, queens, &solutions)

	fmt.Printf("%v solutions in total.\n", len(solutions))
	for _, solution := range solutions {
		fmt.Println(solution)
	}

	var oneSolution []string
	for _, solution := range solutions {
		for y := 0; y < n; y++ {
			row := ""
			for x := 0; x < n; x++ {
				if x == solution[y][0] && y == solution[y][1] {
					row = row + "Q"
				} else {
					row = row + "."
				}
			}
			oneSolution = append(oneSolution, row)
		}
		result = append(result, oneSolution)
		oneSolution = nil
	}
	return result
}

func place(queenNumber int, queens [][2]int, solutions *[][][2]int) {
	y := len(queens)     //from 0 row
	if y < queenNumber { //till n queen placed
		y := len(queens)
		nextRow := make([]bool, queenNumber)
		for x := 0; x < queenNumber; x++ {
			for _, queen := range queens {
				if x-queen[0] == y-queen[1] || //与该 queen 右斜冲突
					x == queen[0] || //纵向冲突
					x+y == queen[0]+queen[1] { //左斜冲突
					nextRow[x] = true
					break
				}
			}
		}

		//after checking one row:
		for x := 0; x < queenNumber; x++ {
			if !nextRow[x] { //position x not attacked
				new_queens := append(queens, [2]int{x, y})
				place(queenNumber, new_queens, solutions) //check next row
				new_queens = new_queens[:y]               //Reverse try next x column
			}
		}
	} else {
		*solutions = append(*solutions, queens)
	}
}

func main() {
	fmt.Println(solveNQueens(7))
}
