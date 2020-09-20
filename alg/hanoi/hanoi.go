package main

import "fmt"

func main() {
	hanoi(6, 'A', 'B', 'C')
}

//x, y, z three pillars
func hanoi(n int, x, y, z byte) {
	if n == 0 {
	} else {
		hanoi(n-1, x, z, y)
		fmt.Printf("%c -> %c, ", x, y)
		hanoi(n-1, z, y, x)
	}
}
