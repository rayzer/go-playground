package main

import (
	"fmt"
)

func main() {
	var a [4]int
	a[0] = 1
	i := a[0]
	// i == 1
	fmt.Printf("%T %#v\n", i, i)

	//b := [2]string{"Penn", "Teller"}
	// b := [...]string{"Penn", "Teller"}

	//letters := []string{"a", "b", "c", "d"}

	//var s []byte
	//s = make([]byte, 5, 5)
	// s = []byte{0, 0, 0, 0, 0}
	s := make([]byte, 5, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	//b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//b[1:4] == []byte{'o', 'l', 'a'}
	// b[:2] == []byte{'g', 'o'}
	// b[2:] == []byte{'l', 'a', 'n', 'g'}
	// b[:] == b

	//x := [3]string{"Лайка", "Белка", "Стрелка"}
	//t := x[:]

	d := []byte{'r', 'o', 'a', 'd'}

	fmt.Printf("d: %#v, %v\n", d, d)

	e := d[2:]
	// e == []byte{'a', 'd'}
	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}
	fmt.Printf("e: %#v, %v\n", e, e)
	fmt.Printf("d: %#v, %v\n", d, d)

	s = s[2:4]
	s = s[:cap(s)]

	//Double s capacity

	// t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
	// for i := range s {
	// 	t[i] = s[i]
	// }
	// s = t

	t := make([]byte, len(s), (cap(s)+1)*2)
	copy(t, s)
	s = t

	//append
	as := make([]int, 1)
	// a == []int{0}
	fmt.Printf("as: %#v, %v\n", as, as)

	as = append(as, 1, 2, 3)
	// a == []int{0, 1, 2, 3}
	fmt.Printf("as: %#v, %v\n", as, as)

	ad := []string{"John", "Paul"}
	printMyslices(ad)
	bd := []string{"George", "Ringo", "Pete"}
	printMyslices(bd)
	ad = append(ad, ad...) // equivalent to "append(a, b[0], b[1], b[2])"
	// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
	printMyslices(ad)
}

func printMyslice(a []int) {
	fmt.Printf("%#v, %v\n", a, a)
}

func printMyslices(a []string) {
	fmt.Printf("%#v, %v\n", a, a)
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
