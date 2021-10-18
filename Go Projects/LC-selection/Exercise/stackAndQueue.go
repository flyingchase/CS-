package Exercise

import "math"

type MinStack struct {
	min   []int
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

// 使用栈实现队列
type MyQueue struct {
	// 双栈模拟队列
	// stack 栈为模拟入队操作
	// 	入队前即 Push 操作需要先将 back 带出队元素全部纳入入队栈内(FIFO)
	// back 栈为模拟出队操作
	// 	出队前即 Pop 需要将 stack 入队元素从栈顶抛出加入 back 栈上方(保证先进先出)
	stack []int
	back []int
}

func ConstructorStack()MyQueue  {
	return MyQueue{
		stack: make([]int,0),
		back: make([]int,0),
	}
}
func (this *MyQueue)Push(x int)  {
	// 从 back 栈中依次弹出存储到 stack 栈中
	// x 存储到 stack 栈栈顶
	for len(this.back)!=0 {
		val:=this.back[len(this.back)-1]
		this.back=this.back[:len(this.back)-1]
		this.stack=append(this.stack,val)
	}
	this.stack=append(this.stack,x)
}

func (this *MyQueue)Pop()int  {
	// 弹出队头元素：
	// 	back 栈栈顶元素为队头，
	for len(this.stack)!=0 {
		val:=this.stack[len(this.stack)-1]
		this.stack=this.stack[:len(this.stack)-1]
		this.back=append(this.back,val)
	}
	// 双栈均空
	if len(this.back)==0 {
		return 0
	}
	val:=this.back[len(this.back)-1]
	this.back=this.back[:len(this.back)-1]
	return val
}

func (this *MyQueue) Peek() int  {
	// 先保证所有元素都在出队栈内
	for len(this.stack)!=0 {
		val:=this.stack[len(this.stack)-1]
		this.stack=this.stack[:len(this.stack)-1]
		this.back=append(this.back,val)
	}
	if len(this.back)==0 {
		return 0
	}
	// 无须将出队栈栈顶元素抛去
	val:=this.back[len(this.back)-1]
	return val
}
func (this *MyQueue) Empty()bool  {
	return len(this.stack)==0&&len(this.back)==0
}
