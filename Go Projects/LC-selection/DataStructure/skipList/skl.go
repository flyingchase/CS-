package skipList

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

/*
实现跳表 即跳跃链表
	- 在 logn 时间复杂度下插入、删除、查找操作


*/

const (
	maxLevel    int     = 18
	probability float64 = 1 / math.E
)

type handleEle func(e *Element) bool
type (
	Node struct {
		next []*Element
	}
	Element struct {
		Node
		key   []byte
		value interface{}
	}

	SkipList struct {
		Node
		maxLevel       int
		Len            int
		randSource     rand.Source
		probability    float64
		probTable      []float64
		prevNodesCache []*Node
	}
)

func NewSkipList() *SkipList {
	return &SkipList{
		Node: Node{
			next: make([]*Element, maxLevel),
		},
		maxLevel:       maxLevel,
		probability:    probability,
		randSource:     rand.New(rand.NewSource(time.Now().UnixNano())),
		probTable:      probabilityTable(probability, maxLevel),
		prevNodesCache: make([]*Node, maxLevel),
	}

}

func probabilityTable(probablity float64, maxLevel int) (table []float64) {
	for i := 1; i <= maxLevel; i++ {

		prob := math.Pow(probablity, float64(i-1))
		table = append(table, prob)
	}

	return table

}
func (e *Element) Key() []byte {
	return e.key
}
func (e *Element) Value() interface{} {
	return e.value
}
func (e *Element) SetValue(value interface{}) {
	e.value = value
}

func (e *Element) Next() *Element {
	return e.next[0]
}

func (t *SkipList) Front() *Element {
	return t.next[0]
}
func (t *SkipList) Put(key []byte, value interface{}) *Element {
	var element *Element
	prev := t.backNodes(key)
	if element = prev[0].next[0]; element != nil && bytes.Compare(element.key, key) <= 0 {
		element.value = value
		return element
	}
	element = &Element{
		Node:  Node{next: make([]*Element, t.randomLevel())},
		key:   key,
		value: value,
	}

	for i := range element.next {
		element.next[i] = prev[i].next[i]
		prev[i].next[i] = element
	}

	t.Len++
	return element
}

// 找到key 对应的前一个节点的索引
func (t *SkipList) backNodes(key []byte) []*Node {
	var prev = &t.Node
	var next *Element
	prevs := t.prevNodesCache
	for i := t.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && bytes.Compare(key, next.key) > 0 {
			prev = &next.Node
			next = next.next[i]
		}
		prevs[i] = prev
	}
	return prevs
}

// 生成索引的随机层数
func (t *SkipList) randomLevel() (level int) {
	r := float64(t.randSource.Int63()) / (1 << 63)

	level = 1
	for level < t.maxLevel && r < t.probTable[level] {
		level++
	}
	return
}

func (t *SkipList) Get(key []byte) *Element {
	var prev = &t.Node
	var next *Element

	for i := t.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && bytes.Compare(key, next.key) > 0 {
			prev = &next.Node
			next = next.next[i]
		}
	}
	if next != nil && bytes.Compare(next.key, key) <= 0 {
		return next
	}
	return nil
}

func (t *SkipList) Exit(key []byte) bool {
	return t.Get(key) != nil
}

func (t *SkipList) Remove(key []byte) *Element {
	prev := t.backNodes(key)
	if element := prev[0].next[0]; element != nil && bytes.Compare(element.key, key) <= 0 {
		for k, v := range element.next {
			prev[k].next[k] = v
		}
		t.Len--
		return element
	}
	return nil
}

func (t *SkipList) Foreach(fun handleEle) {
	for p := t.Front(); p != nil; p = p.next {
		if ok := fun(p); !ok {
			break
		}

	}
}

func (t *SkipList) FindPrefix(prefix []byte) *Element {
	var prev = &t.Node
	var next *Element
	for i := t.maxLevel - 1; i >= 0; i-- {
		next = prev.next[i]
		for next != nil && btes.Compare(prefix, next.key) > 0 {
			prev = &next.Node
			next = next.next[i]
		}
	}
	if next == nil {
		next = t.Front
	}
	return next
}
