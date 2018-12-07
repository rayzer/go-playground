package main

import (
	"fmt"
	"reflect"
)

type data struct {
	num    int
	checks [10]func() bool   // 无法比较
	doIt   func() bool       // 无法比较
	m      map[string]string // 无法比较
	bytes  []byte            // 无法比较
}

func main() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2)) // true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(m1, m2)) // true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// 注意两个 slice 相等，值和顺序必须一致
	fmt.Println("v1 == v2: ", reflect.DeepEqual(s1, s2)) // true
}
