package main

import "fmt"

func solveNQueens(n int) [][]string {
	var queens [][2]int
	var result [][]string
	var solutions [][][2]int
	place(n, queens, &solutions)

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

		for x := 0; x < queenNumber; x++ {
			if !nextRow[x] { //position x not attacked
				queens = append(queens, [2]int{x, y}) //place a queen on x, y
				place(queenNumber, queens, solutions) //check next row (if there is a solution, will be stored)
				queens = queens[:y]                   //try next x column
			}
		}
	} else {
		*solutions = append(*solutions, queens)
	}
}

//func main() {
//	var queens [][2]int
//	var result [][]string
//	place(queens)
//	fmt.Println("solutions found: ", solutionCount)
//	var oneSolution []string
//	for _, solution := range solutions {
//		for y := 0; y < n; y++ {
//			row := ""
//			for x := 0; x < n; x++ {
//				if x == solution[y][0] && y == solution[y][1] {
//					row = row + "Q"
//				} else {
//					row = row + "."
//				}
//			}
//			oneSolution = append(oneSolution, row)
//		}
//		result = append(result, oneSolution)
//		oneSolution = []string{}
//	}
//	fmt.Println(result)
//}
func main() {
	fmt.Println(solveNQueens(6))
}
