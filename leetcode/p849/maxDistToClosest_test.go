//https://leetcode.com/contest/weekly-contest-88/problems/maximize-distance-to-closest-person/

package p849

import (
	"log"
	"testing"
)

func maxDistToClosest(seats []int) int {
	numberOfSeats := len(seats)
	anwser, before, after := 0, -1, 0

	for i := 0; i < numberOfSeats; i++ {
		if seats[i] == 1 {
			before = i
		} else {
			//goes to the end of 0s
			for (after < numberOfSeats && seats[after] == 0) || after < i {
				after++
			}

			var left, right int
			if before == -1 {
				left = numberOfSeats
			} else {
				left = i - before
			}

			if after == numberOfSeats {
				right = numberOfSeats
			} else {
				right = after - i
			}
			log.Println("position:", i, "left:", left, "right:", right)
			anwser = max(anwser, min(left, right))
		}
	}
	return anwser
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test(t *testing.T) {
	if maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}) != 2 {
		t.Fail()
	}
}
