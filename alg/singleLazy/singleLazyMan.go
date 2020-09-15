package main

import (
	"fmt"
	"sync"
)

type singleLazyMan struct {
	who string
}

var (
	instance *singleLazyMan
	once     sync.Once
)

func GetLazyMan() *singleLazyMan {
	once.Do(func() {
		instance = &singleLazyMan{"you"}
	})
	return instance
}

func main() {
	a := GetLazyMan()
	fmt.Println(a.who)
}
