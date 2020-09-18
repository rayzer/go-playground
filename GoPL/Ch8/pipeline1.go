package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//不停往 natuals 刷x
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//不停取出
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
	}
}
