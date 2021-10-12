package WarmUp

type Node struct {
	prev,next *Node
	key int
	value int
	lists LRUCache
}

type LRUCache struct {
	cap int
	root *Node
	len int
}


func Constructor(capacity int) LRUCache {
	node:=&Node{}
	node.prev,node.next=node,node
	l:=LRUCache{
		cap : capacity,
		root : node,
	}
	l.root.lists=l

	return l
}


func (this *LRUCache) Get(key int) int {
	n:=this.get(key)
	if n==nil {
		return -1
	}
	return n.value
}
// 每次 get 不存在返回 -1
// 存在返回 value 并将 node 移动至链首
func (l *LRUCache)get (key int) *Node {
	for n:=l.root.next;n!=l.root;n=n.next {
		if n.key==key {
			n.prev.next=n.next
			n.next.prev=n.prev


			n.prev=l.root
			n.next=l.root.next
			l.root.next.prev=n
			l.root.next=n
			return n
		}
	}
	return nil

}

func (this *LRUCache) Put(key int, value int)  {
	this.len++
	if this.len>this.cap {
		last:=this.root.prev
		last.prev.next=this.root
		this.root.prev=last.prev

		last.next,last.prev=nil, nil
	}
	n:=&Node{
		value: value,
		key: key,
	}
	n.prev=this.root
	n.next=this.root.next
	this.root.next.prev=n
	this.root.next=n

	n.lists=this


}