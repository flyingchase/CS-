package Dailyguoguo

import (
	"LC-selection/DataStructure"
)

type TreeNode = DataStructure.TreeNode

var m map[int]int
var mostFreq int

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func findSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val + findSum(root.Left) + findSum(root.Right)
	m[sum]++
	if m[sum] > mostFreq {
		mostFreq = m[sum]
	}
	return sum
}
func findFrequentTreeSum(root *TreeNode) []int {

	m = make(map[int]int)
	mostFreq = 0
	res := []int{}
	findSum(root)
	for k, v := range m {
		if v == mostFreq {
			res = append(res, k)
		}
	}
	return res
}
