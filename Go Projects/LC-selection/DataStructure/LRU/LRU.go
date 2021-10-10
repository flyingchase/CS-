package LRU

import "fmt"

// LRU represent Least Recently Used strategy
// means
// cap 固定的双向链表
// 每次插入均在头部,已经存在的元素则更新元素结点并移至头部,不存在则新建结点插入,超容则删掉最后一个结点
// 每次查询则移动待查元素到头部(更新使用频率) 不存在
type Node struct {
	prev, next *Node
	// 代表当前 LRU 的归属
	list *LRU

	key   string
	value interface{}
}
type LRU struct {
	root *Node
	cap  int
	len  int
}

func NewLRU(cap int) *LRU {
	l := &LRU{
		root: &Node{},
		cap:  cap,
	}
	l.root.prev = l.root
	l.root.next = l.root
	l.root.list = l
	return l
}

// Get  获得缓存数据
// 获取到则将该节点移动至链表头部
// 为获取到则 nil
func (l *LRU) Get(key string) interface{} {
	defer l.debug()
	n := l.get(key)
	if n == nil {
		return nil
	}
	return n.value
}

func (l *LRU) get(key string) *Node {
	for n := l.root.next; n != l.root; n = n.next {
		if n.key == key {
			n.prev.next = n.next
			n.next.prev = n.prev

			n.next = l.root.next
			l.root.next.prev = n
			l.root.next = n
			n.prev = l.root
			return n
		}
	}
	return nil
}

// 将 key 构造为结点插入头部,若存在对应的 key 则更新结点值
// 缓存满则删掉最后结点(LRU)最少使用

func (l *LRU) Put(key string, value interface{}) {
	defer l.debug()
	n := l.get(key)
	if n != nil {
		n.value = value
		return
	}
	// 缓存 cap 满
	// delete the last node and l.len--
	if l.len == l.cap {
		last := l.root.prev
		last.prev.next = l.root
		l.root.prev = last.prev
		last.prev = nil
		last.next = nil
		last.list = nil
		l.len--
	}
	// construct newNode to insert the head of the list
	// update the l.len and newNode'list
	node := &Node{
		key:   key,
		value: value,
	}
	head := l.root.next
	head.prev = node
	node.next = head
	node.prev = l.root
	l.root.next = node
	l.len++
	node.list = l
}
func (l *LRU) debug() {
	fmt.Println("lru len: ", l.len)

	fmt.Println("lru cap: ", l.cap)
	for n := l.root.next; n != l.root; n = n.next {
		fmt.Printf("%s:%v -> ", n.key, n.value)
	}
	fmt.Println()
}
