package binarySearch

func binsearch(a []int, len int, targetValue int) int {
	low := 0
	high := len - 1

	for low <= high {
		mid := low + (high-low)/2
		if a[mid] == targetValue {
			return mid
		} else if a[mid] < targetValue {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
