package LFU

import (
	"container/list"
)

type (
	node struct {
		key,value int
		frequency int
	}
	LFUCacheWithInnerList struct {
		nodes map[int]*list.Element
		lists map[int]*list.List
		cap int
		min int
	}
)

func Constructor (capacity int) LFUCacheWithInnerList  {
	return LFUCacheWithInnerList{
		nodes: make(map[int]*list.Element),
		lists: make(map[int]*list.List),
		cap: capacity,
		min: 0,
	}
}
// get 需要更新 ferquency 和 nodes lists 两个 map,
// 通过 nodes 的 key 获取节点,在 lists 删除当前 frequency 的结点再++frequency
// 新的 frequency 存在lists 中则添加到链首,不存在则新建链表并移至表首
// 更新结点作为 value 的 nodes
// 更新 min
// 判断原 frequency 对应的链表是否空 空则 min++
func ( this *LFUCacheWithInnerList) Get(key int) int  {
	value,ok:=this.nodes[key]
	if !ok {
		return -1
	}
	curNode:=value.Value.(*node)
	this.lists[curNode.frequency].Remove(value)
	curNode.frequency++
	if _, ok :=this.lists[curNode.frequency];!ok {
		this.lists[curNode.frequency]=list.New()
	}
	newList:=this.lists[curNode.frequency]
	newNode:=newList.PushFront(curNode)
	this.nodes[key]=newNode
	if curNode.frequency-1==this.min&&this.lists[curNode.frequency-1].Len()==0 {
		this.min++
	}
	return curNode.value
}

func (this *LFUCacheWithInnerList) Put(key int,value int)  {
	if this.cap==0 {
		return
	}
	if curValue,ok:=this.nodes[key];ok {
		curNode:=curValue.Value.(*node)
		curNode.value=value
		this.Get(key)
		return
	}
	if this.cap==len(this.nodes) {
		curList:=this.lists[this.min]
		backNode:=curList.Back()
		delete(this.nodes,backNode.Value.(* node).key)
		curList.Remove(backNode)
	}
	this.min=1
	curNode:=&node{
		key: key,
		value: value,
		frequency: 1,
	}

	if _, ok :=this.lists[1];!ok {
		this.lists[1]=list.New()
	}
	newLlist:=this.lists[1]
	newNode:=newLlist.PushFront(curNode)
	this.nodes[key]=newNode
}