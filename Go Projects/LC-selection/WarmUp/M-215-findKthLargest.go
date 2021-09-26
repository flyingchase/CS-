package WarmUp

import (
	"container/heap"
	"sort"
)

// #1 使用归并排序
func findKthLargest(nums []int, k int) int {
	heapsort215(nums)
	return nums[len(nums)-k]
}

func heapsort215(nums []int) {
	if len(nums)==0 {
		return

	}
	for i,_ := range nums {
		heapInsert215(nums,i)
	}
	size:=len(nums)
	for size > 0 {
		size--
		nums[0],nums[size]=nums[size],nums[0]
		heapIfy215(nums,0,size)
	}

}

func heapIfy215(nums []int, index int, size int) {
	left,right:=index*2+1,index*2+2
	for left<size {
		largest:=left
		if right<size&&nums[left]<nums[right] {
			largest=right
		}
		if nums[index]>nums[largest] {
			largest=index
			break
		}

		nums[index],nums[largest]=nums[largest],nums[index]
		index=largest
		left=2*index+1
		right=2*index+2
	}

}

func heapInsert215(nums []int, index int) {
	parent:=(index-1)/2
	for parent>=0&&nums[parent]<nums[index] {
		nums[parent],nums[index]=nums[index],nums[parent]
		index=parent
		parent=(index-1)>>1
	}
}
// #2 使用堆这一数据结构
type IntHeap []int

func (h IntHeap)Len() int  {
	return len(h)
}
func (h IntHeap) Swap(i,j int)   {
	h[i], h[j]=h[j],h[i]
}
func (h IntHeap) Less (i,j int) bool  {
	return h[i]<h[j]
}
func (h *IntHeap)Push(x interface{})  {
	*h=append(*h,x.(int))
}
func (h *IntHeap) Pop() (val interface{})  {
	old:=*h
	val=old[:len(old) - 1 ]
	*h=old[:len(old) - 1]
	return
}


func findKthLargest2(nums []int, k int) int {
	if len(nums)<k {
	return 0
	}
	h:=&IntHeap{}
	heap.Init(h)
	for i:=0;i<k; i++  {
		heap.Push(h,nums[i])

	}
	for i:=k;i<len(nums);i++ {
		if nums[i]<(*h)[0] {
			continue
		}
		heap.Pop(h)
		heap.Push(h,nums[i])
	}
	return (*h)[0]
}

// #3 sort 直接对 nums 切片进行排序
func findKthLargest3(nums []int, k int) int {
	sort.Slice(nums,func(i, j int) bool{
		return nums[i]>nums[j]
	})
	return nums[k-1]

}