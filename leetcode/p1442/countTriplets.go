package main

func countTriplets(arr []int) int {
	var n = len(arr)
	var result, prefix, c, t int
	var count = make(map[int]int)
	var total = make(map[int]int)
	count[0] = 1
	for i := 0; i < n; i++ {
		prefix = prefix ^ arr[i]
		c = count[prefix]
		t = total[prefix]
		result += c*i - t
		count[prefix] = c + 1
		total[prefix] = t + i + 1
	}
	return result
}
