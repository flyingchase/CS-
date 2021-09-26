package singlyLinkedlist

import "fmt"

/*
实现单链表
	- append
	- 结点前后插入
	- 依值查找
	- 删除结点
	- 打印链表值
*/

type ListNode struct {
	Val  interface{}
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
	size int
}

// 创建结点
func newNode(val interface{}) *ListNode {
	return &ListNode{

		Val:  val,
		Next: nil,
	}
}
func New() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

//添加至链尾 并返回生成的结点
func (lists *LinkedList) PushBack(val interface{}) *ListNode {
	head := lists.head
	// 将待插入值转化为 node
	node := newNode(val)
	if head == nil {
		lists.head = node
	} else {
		for head.Next != nil {
			head = head.Next
		}
		head.Next = node
	}
	lists.size++
	return node
}

//结点前插入
func (lists *LinkedList) PushBefore(p *ListNode, val interface{}) *ListNode {
	if p == nil || lists.head == nil {
		return nil
	}
	node := newNode(val)
	// 在头结点之前插入情况
	if p == lists.head {
		node.Next = lists.head
		lists.head = node
	} else {
		prev := lists.head
		// 找到待插入节点的前一个节点
		for ; prev.Next != p; prev = prev.Next {
		}
		node.Next = p
		prev.Next = node
	}
	lists.size++
	return node
}

//结点后插入
func (lists *LinkedList) PushAfter(p *ListNode, val interface{}) *ListNode {

	if p == nil {
		return nil
	}
	node := newNode(val)
	node.Next = p.Next
	p.Next = node
	lists.size++
	return node
}

//查找结点
func (lists *LinkedList) Find(val interface{}) *ListNode {
	cur := lists.head
	for cur != nil && cur.Val != val {
		cur = cur.Next
	}
	return cur
}

//删除指定结点
func (lists *LinkedList) Delete(p *ListNode) {
	if p == nil {
		return
	}
	if p == lists.head {
		// 待删结点为头结点
		lists.head = lists.head.Next
		lists.size--
	} else {
		// 非删除头结点时 找到待删结点的 prev 结点
		prev := lists.head
		for prev != nil && prev.Next != p {
			prev = prev.Next
		}
		// 保证找到 而非 prev 走到链尾
		if prev != nil {
			prev.Next = p.Next
			lists.size--
		}
	}
}

//删除指定值结点 调用 find 找到指定值的结点再删除结点即可
func (lists *LinkedList)DeleteVal(val interface{})  {
	lists.Delete(lists.Find(val))
}
//打印链表值
func (lists *LinkedList)PrintDara()  {
	if lists.size==0 {
		return
	}
	for node:=lists.head;node!=nil;node = node.Next {
		fmt.Print(node.Val," ")
	}
	fmt.Println()
}
