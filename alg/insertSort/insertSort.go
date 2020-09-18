package main

import "fmt"

func main() {
	notOrdered := []int{5, 1, 0, 2, 3, 4}
	insertSort(notOrdered, len(notOrdered))
	fmt.Println(notOrdered)
}

func insertSort(a []int, len int) {
	if len <= 1 {
		return
	}

	for i := 1; i < len; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}
