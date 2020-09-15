package main

import (
	"fmt"
	"sort"
)

func main() {
	example := []int{1, 3, 4, 2, 7, 8, -9}
	qSort(example)
	fmt.Println(sort.IntsAreSorted(example))
	fmt.Printf("%v\n", example)
	fmt.Printf("%#v\n", example)
	fmt.Printf("%T\n", example)
}

func qSort(data []int) {
	qSort_c(data, 0, len(data)-1)
}

func qSort_c(data []int, p int, r int) {
	if p >= r {
		return
	}

	pivot := partition(data, p, r)
	qSort_c(data, p, pivot-1)
	qSort_c(data, pivot+1, r)
}

func partition(data []int, p int, r int) int {
	pivot := data[r]
	i := p
	for j := p; j < r-1; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[r] = data[r], data[i]
	return i
}

//func qSort(data []int) {
//	qSort_c(data, 0, len(data)-1)
//}
//
//func qSort_c(data []int, p, r int) {
//	if p > r {
//		return
//	}
//	q := partition(data, p, r)
//	qSort_c(data, p, q-1)
//	qSort_c(data, q+1, r)
//}
//
//func partition(data []int, p int, r int) int {
//	pivot := data[r]
//	i := p
//	for j := p; j < r-1; j++ {
//		if data[j] < pivot {
//			data[i], data[j] = data[j], data[i]
//			i++
//		}
//	}
//	data[i], data[r] = data[r], data[i]
//	return i
//}
