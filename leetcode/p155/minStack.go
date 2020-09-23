package main

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start

//16 ms, faster than 90.73%
type MinStack struct {
	Data     []int
	MinStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if len(this.MinStack) == 0 || x <= this.GetMin() {
		this.MinStack = append(this.MinStack, x)
	}
}

func (this *MinStack) Pop() {
	if this.Data[len(this.Data)-1] == this.GetMin() {
		this.MinStack = this.MinStack[:len(this.MinStack)-1]
	}
	this.Data = this.Data[:len(this.Data)-1]

}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
