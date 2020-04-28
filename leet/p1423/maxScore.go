package main

func maxScore(cardPoints []int, k int) int {
	var maxScore int
	for i := 0; i < k; i++ {
		maxScore += cardPoints[i]
	}
	var tmpSum = maxScore
	for i := k - 1; i >= 0; i-- {
		tmpSum -= cardPoints[i]
		tmpSum += cardPoints[len(cardPoints)-(k-i)]
		if tmpSum > maxScore {
			maxScore = tmpSum
		}
	}
	return maxScore
}
