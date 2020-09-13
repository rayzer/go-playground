package main

func buildArray(target []int, n int) []string {
	var result []string
	var targetIdx = 0
	for i := 1; i <= n; i++ {
		result = append(result, "Push")
		if target[targetIdx] == i {
			targetIdx++
			if targetIdx >= len(target) {
				break
			}
			continue
		} else {
			result = append(result, "Pop")
		}
	}
	return result
}
