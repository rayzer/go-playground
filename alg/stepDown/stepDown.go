package main

// 1: 1
// 2: 11, 2
// 3: 111, 12, 21
// f(3) = f(2） + f(1)
//f(n) = f(n-1) + f(n-2)

func main() {
	println(stepDown(10))
}

func stepDown(n int) int {
	count := 0
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n > 2 {
		for i := 3; i <= n; i++ {
			count = stepDown(n-1) + stepDown(n-2)
		}
	}
	return count
}
