package main

import (
	"fmt"
	"sync"
)

type instance struct {
	Values map[string]string
}

var (
	once sync.Once
	i    *instance
	s    *sman
)

func new() *instance {
	once.Do(func() {
		i = &instance{
			Values: make(map[string]string),
		}
	})
	return i
}

func main() {
	s := new()
	s.Values["al"] = "ar"
	s2 := new()
	s2.Values["al"] = "ar2"
	fmt.Printf("%v : %v ", "al", s2.Values["al"])
}

type sman struct {
	value string
}

func newInstance() *sman {
	once.Do(func() {
		s = &sman{
			value: "ray",
		}
	})
	return s
}
