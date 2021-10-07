package Heap

import (
	"errors"
)

/*
实现二叉堆
	- 添加
	- 删除堆顶
	- 获取堆顶
	- 堆化
*/

const DefaultCapacity = 16

type BinaryHeap struct {
	data     []int
	size     int
	capacity int
}

func Default() *BinaryHeap {
	return New(DefaultCapacity)
}

func New(capacity int) *BinaryHeap {
	return &BinaryHeap{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

// 添加元素
//  现将元素添加到堆尾 再自下而上不断比较 (index-1)/2比较 移至合适位置
func (h *BinaryHeap) Add(val int) bool {
	if h.size >= h.capacity {
		return false
	}
	// 将 val 加入堆尾 再自下向上调整 不断与 parent=(index-1)/2 比较
	h.data[h.size] = val
	h.size++
	index := h.size - 1
	parent := (index - 1) / 2
	for parent >= 0 && h.data[index] < h.data[parent] {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index, parent = parent, (index-1)/2
	}
	return true
}

// 删除堆顶元素
func (h *BinaryHeap) Remove() error {
	if h.size <= 0 {
		return errors.New("heap is null")
	}
	h.size--
	if h.size > 0 {
		h.data[0] = h.data[h.size-1]
		shiftDown(h, 0, h.size)
	}
	return nil

}

func shiftDown(h *BinaryHeap, index int, size int) {
	left, right := 2*index+1, 2*index+2
	largest := left
	for index < size {
		if right < size && h.data[right] < h.data[left] {
			largest = right
		}
		if h.data[index] > h.data[largest] {
			largest = index
			break
		}
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
		left = 2*index + 1
		right = 2*index + 2
	}
}

// 获取堆顶元素

func (h BinaryHeap) GetTop() int {
	if h.size <= 0 {
		panic("heap is null")
	}
	return h.data[0]
}
