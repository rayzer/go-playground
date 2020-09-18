package main

import "fmt"

func main() {
	twoD := make([][]int, 3)
	fmt.Println(twoD)

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 3; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}
	fmt.Println(twoD)

	for _, x := range twoD {
		fmt.Println(x)
		for i, y := range x {
			fmt.Println("i:", i, "value:", y)
		}
	}
}
