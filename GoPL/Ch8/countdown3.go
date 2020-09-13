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

	ticker := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker:
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}

	}
	launch()
}

func launch() {
	fmt.Println("LAUNCH")
}
