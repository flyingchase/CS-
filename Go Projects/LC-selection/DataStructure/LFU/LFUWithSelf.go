package LFU

import (
	"container/heap"
)

type (
	nodeLFU struct {
		key, value int
		// priority in heap
		frequency int
		// count means which is the oldest 操作时间戳
		count int
		index int
	}
	LFUCacheWithSelf struct {
		cap     int
		pq      priorityQueue
		hash    map[int]*nodeLFU
		counter int
	}
)
type priorityQueue []*nodeLFU

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].frequency == pq[j].frequency {
		return pq[i].count < pq[j].count
	}
	return pq[i].frequency < pq[j].frequency
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*nodeLFU)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[:n-1]
	return node
}
func (pq *priorityQueue) update(node *nodeLFU, value int, frequency int, count int) {
	node.value = value
	node.frequency = frequency
	node.count = count
	heap.Fix(pq, node.index)
}

func ConstructorLFUWithSelf(capacity int) LFUCacheWithSelf {
	return LFUCacheWithSelf{
		cap:  capacity,
		hash: make(map[int]*nodeLFU, capacity),
		pq:   priorityQueue{},
	}
}
func (this *LFUCacheWithSelf) Get(key int) int {
	if this.cap == 0 {
		return -1
	}
	if node, ok := this.hash[key]; ok {
		this.counter++
		this.pq.update(node, node.value, node.frequency+1, node.count)
		return node.value
	}
	return -1
}
func (this *LFUCacheWithSelf) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	this.counter++
	if node, ok := this.hash[key]; ok {
		this.pq.update(node, node.value, node.frequency+1, node.count)
		return
	}
	if len(this.pq) == this.cap {
		node := heap.Pop(&this.pq).(*nodeLFU)
		delete(this.hash, node.key)
	}
	node := &nodeLFU{
		value: value,
		key:   key,
		count: this.counter,
	}
	heap.Push(&this.pq, node)
	this.hash[key] = node
}
