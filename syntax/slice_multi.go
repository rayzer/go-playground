package main

import "fmt"

func main() {
	result := make([][]int, 0)
	for i := 0; i < 3; i++ {
		result = append(result, []int{})
		result[i] = make([]int, 0, 3)
		result[i] = append(result[i], i)
		result[i] = append(result[i], i)
		result[i] = append(result[i], i)
		result[i] = append(result[i], i)
		result[i] = append(result[i], 1)
	}
	fmt.Println(result)
}
