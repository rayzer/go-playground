package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("1")
	case <-abort:
		fmt.Println("Launch aborted")
		return
	}
	launch()
}

func launch() {
	fmt.Println("LAUNCH")
}
