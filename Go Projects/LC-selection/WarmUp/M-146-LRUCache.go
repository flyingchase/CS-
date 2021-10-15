package WarmUp


// type Node struct {
// 	prev, next *Node
// 	key        int
// 	value      int
// 	lists      *LRUCache
// }
//
// type LRUCache struct {
// 	cap  int
// 	root *Node
// 	len  int
// }
//
// func Constructor(capacity int) LRUCache {
// 	node := &Node{}
// 	node.prev, node.next = node, node
// 	l := LRUCache{
// 		cap:  capacity,
// 		root: node,
// 	}
// 	l.root.lists = &l
// 	return l
// }
//
// func (this *LRUCache) Get(key int) int {
// 	n := this.get(key)
// 	// 每次 get 不存在返回 -1
// 	if n == nil {
// 		return -1
// 	}
// 	return n.value
// }
//
// // 存在返回 value 并将 node 移动至链首
// func (l *LRUCache) get(key int) *Node {
// 	for n := l.root.next; n != l.root; n = n.next {
// 		if n.key == key {
// 			// 截断 n 的上下结点
// 			n.prev.next = n.next
// 			n.next.prev = n.prev
// 			// 将 n查到 root 节点的后面
// 			n.prev = l.root
// 			n.next = l.root.next
// 			l.root.next.prev = n
// 			l.root.next = n
// 			return n
// 		}
// 	}
// 	return nil
//
// }
//
// // 插入节点三种情况
// // - 节点 key 存在则更新节点值
// // - 节点不存在且
// // 	- 插入后超容,先删除最后结点
// // 	- 后插入
// func (this *LRUCache) Put(key int, value int) {
// 	node := this.get(key)
// 	if node != nil {
// 		node.value = value
// 		return
// 	}
// 	// 插入节点后导致超容 删除最后结点
// 	if this.len == this.cap {
// 		last := this.root.prev
// 		last.prev.next = this.root
// 		this.root.prev = last.prev
//
// 		last.next, last.prev = nil, nil
// 		this.len--
// 	}
// 	n := &Node{
// 		value: value,
// 		key:   key,
// 	}
// 	n.prev = this.root
// 	n.next = this.root.next
// 	this.root.next.prev = n
// 	this.root.next = n
// 	n.lists = this
// 	this.len++
// }
// solution2 respect to cookbook leetcode

type (
	Node struct {
		key,val int
		prev,next *Node
	}
	LRUCache struct {
		head,tail *Node
		keys map[int]*Node
		Cap int
	}
)

func Construtor(capacity int) LRUCache  {
	return LRUCache{
		keys: make(map[int]*Node),
		Cap: capacity,
	}
}
func (this *LRUCache)Get(key int) int  {
	if Node,ok := this.keys[key];ok {
		this.remove(Node)
		this.add(Node)
	return Node.val
	}
	return -1

}
func (this *LRUCache)Put (key int,value int)  {
	if node,ok :=this.keys[key];ok {
		node.val=value
		this.remove(node)
		this.add(node)
		return
	}else {
	    node=&Node{
	    	key: key,
	    	val: value,
		}
		this.keys[key]=node
		this.add(node)
	}
	if len(this.keys)>this.Cap {
		delete(this.keys,this.tail.key)
		this.remove(this.tail)
	}
}
func (l *LRUCache)add(node *Node)  {
	node.prev,node.next=nil,l.head
	if l.head!=nil {
		l.head.prev=node
	}
	l.head=node
	if l.tail==nil {
		l.tail=node
		l.tail.next = nil
	}
}
func (l *LRUCache)remove(node *Node)  {
	if node==l.head {
		l.head=node.next
		node.next=nil
		return
	}
	if node==l.tail {
		l.tail=node.prev
		node.prev.next=nil
		node.prev=nil
		return
	}
	node.prev.next=node.next
	node.next.prev=node.prev

}