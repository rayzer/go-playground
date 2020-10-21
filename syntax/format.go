package main

import (
	"fmt"
	"strconv"
)

func main() {
	//binary
	n := int64(123)
	fmt.Println(strconv.FormatInt(n, 2)) // 1111011
	fmt.Println(strconv.FormatInt(-n, 2))
	fmt.Println(strconv.FormatInt(n&-n, 2)) // 1

	n2, _ := strconv.ParseUint("11110110", 2, 8)
	n2_8 := uint8(n2)
	fmt.Printf("%b, %v\n", n2_8, n2_8)   // 1111011
	fmt.Printf("%b, %v\n", -n2_8, -n2_8) //
	fmt.Printf("%b\n", n2_8&-n2_8)       //
}
