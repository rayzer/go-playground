package main

//s台阶，一次可以下 m 阶
//f(s, m) = f(s-1) + f(s-2) + ... + f(s-m)

//f(0, 0) = 0
//f(1, 1) = 1
//f(2, 2) = 11, 2
//f(3, 3) = 111, 21, 12, 3
//f(4, 4) = f(4, 3) + 1 //一次跨4
//f(m, m) = f(m, m-1) + 1

func main() {
	println(stepDown(10, 1))
}

func stepDown(n int, max int) int {
	count := 0

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	} else if n <= max {
		count = stepDown(n, n-1) + 1
	} else if n > max {
		for i := 1; i <= max; i++ {
			count += stepDown(n-i, max)
		}
	}
	return count
}
