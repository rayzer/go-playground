package main

import (
	"fmt"
)

func main() {
	// ordered := []int{0, 1, 2, 2, 3, 4, 5, 5}
	notOrdered := []int{5, 2, 1, 0, 2, 5, 3, 4}

	length := len(notOrdered)
	duplicate := 0
	for i := 0; i < len(notOrdered)-2; {
		if notOrdered[i] != i {
			if notOrdered[i] == notOrdered[notOrdered[i]] {
				notOrdered[i], notOrdered[length-2+duplicate] = notOrdered[length-2+duplicate], notOrdered[i]
				duplicate++
			} else {
				notOrdered[i], notOrdered[notOrdered[i]] = notOrdered[notOrdered[i]], notOrdered[i]
			}
		} else {
			i++
		}
		fmt.Println(notOrdered)
	}

	fmt.Println(notOrdered)
	insert(notOrdered, length-1)
	insert(notOrdered, length-1)
	fmt.Println(notOrdered)
}

func insert(data []int, idx int) {
	value := data[idx]
	for j := len(data) - 1; j >= data[value]+2; j-- {
		data[j] = data[j-1]
	}
	data[value+1] = value
}
