package Exercise

import (
	"LC-selection/DataStructure"
	"sort"
)

type TreeNode = DataStructure.TreeNode

func postTravelsBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Left
	}
	j := len(res) - 1
	for i := 0; i <= j; i++ {
		j--
		res[i], res[j] = res[i], res[j]
	}
	return res
}

func levelTravelsBT(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}
	cur := root
	queue := make([]*TreeNode, 0)
	queue = append(queue, cur)
	for len(queue) > 0 {
		size := len(queue)
		lists := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			lists = append(lists, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, lists)
	}
	return res
}

type TreePosition struct {
	node     *TreeNode
	position int
}

func verticalTravelsBT(root *TreeNode) [][]int {
	levels := make(map[int]map[int][]int, 0)
	deep := 0
	queue := make([]TreePosition, 0)
	cur := root
	queue = append(queue, TreePosition{cur, 0})
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if node.node.Left != nil {
				queue = append(queue, TreePosition{node.node.Left, node.position - 1})
			}
			if node.node.Right != nil {
				queue = append(queue, TreePosition{node.node.Right, node.position + 1})
			}
			level, ok := levels[node.position]
			if !ok {
				level = make(map[int][]int)
			}
			level[deep] = append(level[deep], node.node.Val)
			levels[node.position] = level
		}
		deep++
	}
	res := make([][]int, 0)
	for i := -1000; i < 1000; i++ {
		levelNode, ok := levels[i]
		if !ok {
			continue
		}
		levelVals := make([]int, 0)
		for j := 0; j <= deep; j++ {
			vals, ok := levelNode[j]
			if !ok {
				continue
			}
			sort.Ints(vals)
			levelVals = append(levelVals, vals...)
		}
		res = append(res, levelVals)
	}
	return res
}
func preOrderTravesalDFS(root *TreeNode) []int {
	res := make([]int, 0)
	dfsToBottom(root, &res)
	return res
}

func dfsToBottom(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfsToBottom(root.Left, res)
	dfsToBottom(root.Right, res)
}
func preOrderTraversalDFS2(root *TreeNode) []int {
	res := dividedAndComquer(root)
	return res
}

func dividedAndComquer(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// divide
	left := dividedAndComquer(root.Left)
	right := dividedAndComquer(root.Right)
	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}
func dividedAndConquerTraversalBT(root *TreeNode) []int {
	res := dividedAndConquer(root)
	return res
}
func dividedAndConquer(root *TreeNode) []int {
	res := make([]int, 0)
	// conditions when return
	if root == nil {
		return res
	}
	// divided
	left := dividedAndConquer(root.Left)
	right := dividedAndConquer(root.Right)
	// conquer
	res = append(res, root.Val)
	// join two slices with ...
	res = append(res, left...)
	res = append(res, right...)
	return res
}
func mergeSort(nums []int) []int {
	if len(nums) <= 0 {
		return nums
	}
	mergeSortHelper(nums, 0, len(nums)-1)
	return nums
}

func mergeSortHelper(nums []int, l int, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	mergeSortHelper(nums, l, mid)
	mergeSortHelper(nums, mid+1, r)
	merge(nums, l, mid, r)
}

func merge(nums []int, l int, mid int, r int) {
	p1, p2, i, helper := l, mid+1, 0, make([]int, 0, r-l+1)
	for p1 <= mid && p2 <= r {
		if nums[p1] <= nums[p2] {
			helper[i] = nums[p1]
			p1++
		} else {
			helper[i] = nums[p2]
			p2++
		}
		i++
	}
	copy(helper[i:], nums[p1:mid+1])
	copy(helper[i:], nums[p2:r+1])
	copy(nums, helper)
	// / use slices merge without copy
	/*
		helper=append(helper,nums[p1:mid+1]...)
		helper=append(helper,nums[p2:r+1]...)
		copy(nums,helper)
	*/
}

// divided application in mergesort
func MergeSort(nums []int) []int {
	return MergeSortHelper(nums)

}

func MergeSortHelper(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// divided into two parts
	mid := len(nums) >> 1
	left := MergeSortHelper(nums[:mid])
	right := MergeSortHelper(nums[mid:])
	// conquer
	res := Merge(left, right)
	return res
}

func Merge(left []int, right []int) (res []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			res = append(res, right[r])
			r++
		} else {
			res = append(res, left[l])
			l++
		}
	}
	// merge the part lasted
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return
}

// use divide in quicksort
func quicksortOld(nums []int) []int  {
	return quicksortOldHelper(nums,0,len(nums)-1)
}

func quicksortOldHelper(nums []int, l int, r int) []int {
	if l>=r {
		return nums
	}
	nums[l],nums[r]=nums[r],nums[l]
	p:=partitionOld(nums,l,r)
	quicksortOldHelper(nums,l,p[0]-1)
	quicksortOldHelper(nums,p[1]+1,r)
	return nums
}

func partitionOld(nums []int, l int, r int) []int {

	less,more:=l-1,r
	for l < more {
		if nums[l]<nums[r] {
			less++
			nums[l], nums[less] = nums[less], nums[l]
			l++
		}else if nums[l] <nums[r] {
			more--
			nums[l],nums[more]=nums[more],nums[l]
		}else {
		    l++
		}
	}
	nums[r],nums[more]=nums[more],nums[r]
	return []int {less+1,more}
}
func QucikSort(nums []int) []int  {
	quickSortHelper(nums,0,len(nums)-1)
	return nums
}

func quickSortHelper(nums []int, l int, r int) {
	if l<r {
		// divide
		pivot:=partition(nums,l,r)
		quickSortHelper(nums,0,pivot-1)
		quickSortHelper(nums,pivot+1,r)
	}
}

func partition(nums []int, l int, r int) int {
	p,i:=nums[r],l
	for j:=l;j<r;j++ {
		if nums[j]<p {
			nums[j],nums[p]=nums[p],nums[j]
			j++
		}
	}
	// pivot update
	nums[i],nums[r]=nums[r],nums[i]
	return i
}
func postorderTraversal(root *TreeNode)[]int  {
	if root==nil {
		return []int{}
	}
	stack:=make([]*TreeNode,0)
	cur:=root
	res:=make([]int,0)
	for cur != nil||len(stack)!=0 {
		for cur!=nil {
			stack=append(stack, cur)
			res = append(res,cur.Val)
			cur=cur.Right
		}
		node:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur=node.Left
	}
	size:=len(res)
	for i:=0;i<size;i++ {
		res=append(res,res[size-i-1])
	}
	return res[size:]
}

























































