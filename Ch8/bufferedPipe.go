package main

import "fmt"

func main() {
	bufferChannel := make(chan int, 3)
	bufferChannel <- 1
	bufferChannel <- 2
	bufferChannel <- 3
	fmt.Println("cap: ", cap(bufferChannel))
	fmt.Println("len: ", len(bufferChannel))

	fmt.Println(<-bufferChannel)
	fmt.Println("len: ", len(bufferChannel))

	bufferChannel <- 4
	fmt.Println(<-bufferChannel)
	fmt.Println(<-bufferChannel)
	fmt.Println(<-bufferChannel)
	//fmt.Println(<-bufferChannel) //NOT OK

	close(bufferChannel)
}
