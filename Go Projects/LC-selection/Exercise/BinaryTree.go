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


