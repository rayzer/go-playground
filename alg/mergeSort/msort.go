package main

import (
	"fmt"
	"sort"
)

func main() {
	example := []int{1, 4, 3, 2, 7, 8, -9}
	fmt.Println(sort.IntsAreSorted(example))
	mSort(example)
	fmt.Println(sort.IntsAreSorted(example))
	fmt.Printf("%v\n", example)
}

func mSort(data []int) {
	//p >= r
	//mergeSort(p...r) = mergeSort(mergeSort(p...q), mergeSort(q+1, r))
	mSort_c(data, 0, len(data)-1)
}

func mSort_c(data []int, p int, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	mSort_c(data, p, q)
	mSort_c(data, q+1, r)
	merge(data, p, q, r)
}

func merge(data []int, p int, q int, r int) {
	fmt.Println("merging data: ", data[p:r+1], p, q, r)
	fmt.Println("current data: ", data)
	tmp := make([]int, (r - p + 1))
	tmp_idx := 0
	i := p
	j := q + 1

	for i <= q && j <= r {
		if data[i] <= data[j] {
			tmp[tmp_idx] = data[i]
			tmp_idx++
			i++
		} else {
			tmp[tmp_idx] = data[j]
			tmp_idx++
			j++
		}
	}

	start, end := i, q
	if j <= r {
		start, end = j, r
	}

	for start <= end {
		tmp[tmp_idx] = data[start]
		tmp_idx++
		start++
	}

	fmt.Printf("%v\n", tmp)
	copy(data[p:r+1], tmp)
}
