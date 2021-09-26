package doublyLinkedlist

import "fmt"

/*
实现双向链表
	- CRUD
	- 表头和表尾的追加
*/
type ListNode struct {
    Val interface{}
    prev,next *ListNode
}
type LinkedList struct {
    head *ListNode
    size int
}

func New()*LinkedList  {
   return & LinkedList {
       head: nil,
       size: 0,
   }
}
// 插入到表头 区分表头是否空
func (lists * LinkedList)PushFront(val interface{}) *ListNode  {
   node:=newNode(val)
	if lists.head!=nil {
		lists.head.prev=node
		node.next=lists.head
	}
	lists.head=node
	lists.size++
	return node
}

func newNode(val interface{}) *ListNode{
	return & ListNode {
	    Val: val,
	    prev: nil,
	    next: nil,
	}
}

// 插入数据到链尾
func (lists * LinkedList) PushBack(val interface{}) *ListNode{
	//  空表则调用链首插入
	if lists.head==nil {
		lists.PushFront(val)
	}
	node:=newNode(val)
	cur:=lists.head
	for cur.next != nil {
		cur=cur.next
	}
	cur.next=node
	node.prev=cur
	lists.size++
	return node
}
//节点后插入数据
func (lists * LinkedList) PushAfter(p *ListNode, val interface{})  *ListNode {

	if p==nil {
		return nil
	}
	// 找到要插入位置的前后结点
	next:=p.next
	node:=newNode(val)
	// 插入 注意判断 next 为空时next.prev 不存在 无须链接 新插入的 node 即为最后一个节点
	node.next=next
	p.next=node
	node.prev=p
	if next!=nil {
		next.prev=node
	}
	lists.size++
	return node
}
//结点前插入数据
func (lists * LinkedList)PushBefore(p *ListNode,val interface{}) *ListNode {
	if p==nil {
		return nil
	}
	node:=newNode(val)
	prev:=p.prev
	// 待插入的链表为空则调用链首插入函数
	if prev==nil {
		lists.PushFront(val)
	}else {
		// 在指定结点前插入需找到prevNode
		p.prev=node
		node.next=p
		prev.next=node
		node.prev=prev
		lists.size++
	}

	return node
}
//删除结点
func (lists * LinkedList) Delete(p *ListNode)  {
	if p == nil|| lists.head==nil {
		return
	}
	// 删除结点为头结点则直接跳过即可
	if p==lists.head {
		lists.head=p.next
	}else {
		//  注意待删结点的 nextNode 为链尾空节点时 无须链接 nextNode.prev=prevNode
		prevNode,nextNode:=p.prev,p.next
		prevNode.next=nextNode
		if nextNode!=nil {
			nextNode.prev=prevNode
		}
	}
	lists.size--

}
//依值查找结点
func (lists * LinkedList)Find(val interface{}) *ListNode  {
	if lists.head==nil{
		return nil
	}
	cur:=lists.head
	for cur!=nil&&cur.Val!=val {
		cur=cur.next
	}
	return cur
}
//打印链表数据
func (lists * LinkedList)PrintDara()  {
	for p:=lists.head;p!=nil;p=p.next {
		fmt.Print(p.Val," ")
	}
	fmt.Println()
}
func (lists * LinkedList)Size() int  {
	return lists.size
}