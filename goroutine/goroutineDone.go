package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{}, 50)
	for i := 0; i < 50; i++ {
		go func() {
			j := rand.Int63n(10)
			time.Sleep(time.Duration(j * 1000000))
			done <- struct{}{}
		}()
	}

	for len(done) < 50 {
		fmt.Println("waiting jobs to be done. current: ", len(done), " of 50...")
		time.Sleep(time.Millisecond)
	}

	fmt.Println("all done")
}
