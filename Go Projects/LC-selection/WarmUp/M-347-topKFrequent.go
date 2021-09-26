package WarmUp

import (
	"container/heap"
	"sort"
)

// #1
/*
	- map存储数组中出现的元素及其频次
	- pair 作为 sort 的数据结构 再 sort.Sort(sort.reverse())
*/
type Pair struct {
	value int
	count int
}
type PairList []Pair

func (p PairList) Less(i, j int) bool {
	return p[i].count < p[j].count
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PairList) Len() int {
	return len(p)
}
func topKFrequent1(nums []int, k int) []int {

	m := map[int]int{}
	for _, v := range nums {
		m[v] += 1
	}

	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))

	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, p[i].value)
	}
	return res
}

/*
#2 利用堆 存储 kv k 为 value v 为 count


*/
type element struct {
	value int
	count int
}
type PriorityQueue []element

func (pq PriorityQueue)Less(i,j int) bool  {
	return pq[i].count>pq[j].count
}
func (pq PriorityQueue) Swap(i,j int)  {
	pq[i],pq[j] = pq[j],pq[i]
}
func (pq PriorityQueue)Len() int {
	return len(pq)
}
func (pq *PriorityQueue) Push(x interface{})  {
	*pq = append(* pq,x.(element))
}
func (pq *PriorityQueue)Pop()interface{}  {
	old := * pq
	top:=old[len(old)-1]
	*pq=old[:len(old)-1]
	return top
}
func topKFrequent(nums []int, k int) []int {
	m:=map[int] int {}
	for _,v:=range nums {
		m[v]+=1

	}
	pq:=&PriorityQueue{}
	heap.Init(pq)
	for k,v:=range m{
		heap.Push(pq,element{k, v})
	}
	res:=make([]int,0)
	for i:=0;i<k;i++{
		res=append(res,heap.Pop(pq).(element).value)
	}
	return res
}