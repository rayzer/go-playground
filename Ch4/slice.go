package main

import "fmt"

//slice stack
stack = append(stack, v)
top := stack[len(stack) - 1]
stack = stack[:len(stack) - 1]

//remove item from slice
func remove(slice []int, i int) []int{
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

//remove by overwrite i by last item
func remove_without_order(slice []int, i int) []int  {
	slice[i]=slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s:= []int{5,6,7,8,9}
	fmt.Println(remove(s,2))
}