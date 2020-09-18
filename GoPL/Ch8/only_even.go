package main

import "fmt"

//ch 容量为1， 每次 select 只有一个情况被满足，第一次放入0 到 ch，第二次取出 ch 到 x 打印，所以打印0，2，4，6，8
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
