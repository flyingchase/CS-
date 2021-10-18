package StackAndQueue

// 数组实现使用 slice 借助 slice 实现自动扩容
type StackWithArray struct {
	items []string
	// cur 代表栈顶指针
	cur int
}

func NewStackWithArray() *StackWithArray {
	return &StackWithArray{
		items: make([]string, 0),
		cur:   0,
	}
}
func (s *StackWithArray) Push(item string) {
	s.cur++
	// items底层 slice 是否满
	if s.cur == len(s.items) {
		s.items = append(s.items, item)
		return
	}
	// 未满则在stack 对应的 cur 加入
	s.items[s.cur] = item
}
func (s *StackWithArray) Pop() string {
	// 空栈则返回零值
	if s.cur == 0 {
		return ""
	}
	item := s.items[s.cur]
	s.cur--
	return item
}
