package SelfStack

import (
	"fmt"
	"strconv"
)

// 实现计算器

// 操作符的优先级
var operatorPriority = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"(": 2,
	")": 2,
}

// 实现 stack 存储 int
type (
	StackInt struct {
		items []int
		cur   int
	}
	Stack struct {
		root *node
		len  int
	}
)

func NewStack() *Stack {
	n := &node{}
	n.prev = n
	n.next = n

	return &Stack{
		root: n,
	}
}
func (s *Stack) Push(val string) {
	n := &node{val: val}
	s.root.prev.next = n
	n.prev = s.root.prev
	n.next = s.root
	s.root.prev = n
	s.len++
}
func (s *Stack) Pop() string {
	n := s.root.prev
	if n == s.root {
		return ""
	}
	s.root.prev = n.prev
	n.prev.next = s.root

	n.prev, n.next = nil, nil
	s.len--
	return n.val
}
func NewStackInt() *StackInt {
	return &StackInt{
		items: make([]int, 0),
	}
}
func (s *StackInt) Push(val int) {
	s.cur++
	if s.cur == len(s.items) {
		s.items = append(s.items, val)
		return
	}
	s.items[s.cur] = val
}
func (s *StackInt) Pop() int {
	if s.cur == 0 {
		return 0
	}
	item := s.items[s.cur]
	s.cur--
	return item
}

type Calculator struct {
	nums      *StackInt
	operators *Stack
	exp       string
}

func NewCalculator(exp string) *Calculator {
	return &Calculator{
		nums:      NewStackInt(),
		operators: NewStack(),
		exp:       exp,
	}
}

func (c *Calculator) Calculate() int {
	l := len(c.exp)
	for i := 0; i < l; i++ {
		switch e := (c.exp[i]); e {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// 继续向后读取
			j := i
			for j < l && c.exp[j] <= '9' && c.exp[j] >= '0' {
				j++
			}
			n, _ := strconv.Atoi(c.exp[i:j])
			i = j - 1
			c.nums.Push(n)
		case '+', '-', '*', '/':
			// from operators stasck 获取栈顶元素
			pre := c.operators.Pop()
			for pre != "" && pre != "(" && operatorPriority[string(e)] <= operatorPriority[pre] {
				c.nums.Push(c.calc(pre))
				pre = c.operators.Pop()
			}
			if pre != "" {
				c.operators.Push(pre)
			}
			c.operators.Push(string(e))
		case '(':
			// 碰到右侧括号之前一致弹出
			for o := c.operators.Pop(); o != "(" && o != ""; o = c.operators.Pop() {
				c.nums.Push(c.calc(o))
			}
		default:
			panic("invalid  exp")
		}
	}
	o := c.operators.Pop()
	if o == "" {
		return c.nums.Pop()
	}
	return c.calc(o)
}

// calc 单次计算的操作 传入 o 为计算符
func (c *Calculator) calc(o string) int {
	b := c.nums.Pop()
	a := c.nums.Pop()

	fmt.Printf("%d %v %d\n", a, o, b)
	switch o {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}
func calculate(s string) int {
	return NewCalculator(s).Calculate()

}
