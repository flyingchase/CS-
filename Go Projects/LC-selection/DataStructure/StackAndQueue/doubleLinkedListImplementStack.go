package SelfStack

// 链式栈 使用双向链表实现
type (
	node struct {
		prev, next *node
		val        string
	}
	StackWithDLists struct {
		// 代表栈所属
		root *node
		len  int
	}
)

func NewStackWithDLists() *StackWithDLists {
	n := &node{}
	n.prev = n
	n.next = n
	return &StackWithDLists{
		root: n,
	}
}
func (s *StackWithDLists) Push(val string) {
	n := &node{val: val}
	// 头插法 node 结点插在 root 之前
	s.root.prev.next = n
	n.prev = s.root.prev
	n.next = s.root
	s.root.prev = n

	s.len++
}

func (s *StackWithDLists) Pop() string {
	item := s.root.prev
	// 栈空弹出零值
	if item == s.root {
		return ""
	}
	s.root.prev = item.prev
	item.prev.next = s.root
	// 待删除结点前后指针置空
	item.next = nil
	item.prev = nil
	s.len--
	return item.val
}
