package main

import (
	"fmt"
	"github.com/rayli/go-playground/Ch6/geometry"
)

func main() {
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance())
}
