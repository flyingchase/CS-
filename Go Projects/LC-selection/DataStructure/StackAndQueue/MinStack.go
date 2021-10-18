package StackAndQueue

import "math"

type MinStack struct {
	// min栈 栈顶存储最小元素,除栈顶外其他元素不管
	min   []int
	// 普通栈
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}
func (this *MinStack) Push(val int) {
	min := this.GetMin()
	if val < min {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min)
	}
	this.stack = append(this.stack, val)
}
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}
func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}
func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return math.MaxInt32
	}
	min := this.min[len(this.min)-1]
	return min
}
