package main

import "fmt"

func print(prev []int, s int, t int) {
	if prev[t] != -1 && t != s {
		print(prev, s, prev[t])
	}
	fmt.Printf("%v ", t)
}

func main() {
	prev := [...]int{-1, 0, 1, 0, 1, 4, 4, -1}
	print(prev[:], 0, 6)
}

//0 6
//0 4
//0 1
//0 0
//0 1 4 6
