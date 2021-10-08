package Dailyguoguo

import (
	"sort"
)

type TreePosition struct {
	// 存储结点和所在列的 position
	node     *TreeNode
	position int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 存储的是[position][level]node.vals 每列每层上结点的值
	levels := make(map[int]map[int][]int)
	queue := make([]TreePosition, 0)
	queue = append(queue, TreePosition{root, 0})
	// 代表树的层深度
	deep := 0
	// BFS 遍历
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.node.Left != nil {
				queue = append(queue, TreePosition{node: curNode.node.Left, position: curNode.position - 1})
			}
			if curNode.node.Right != nil {
				queue = append(queue, TreePosition{node: curNode.node.Right, position: curNode.position + 1})
			}
			// 找到这一列对应的层，不存在则创建
			levelNode, ok := levels[curNode.position]
			if !ok {
				levelNode = make(map[int][]int)
			}
			// 将该结点的 val 存放在对应层后
			levelNode[deep] = append(levelNode[deep], curNode.node.Val)
			// 更新这一层的值
			levels[curNode.position] = levelNode
		}
		// 每层遍历结束下移层数
		deep++
	}

	// 以下为从存储的 level([position][level][]vals)中按列读取
	res := make([][]int, 0)
	for i := -1000; i < 1000; i++ {
		// 从有数据存储的列开始读取，读取出的是每层的 level[level]vals ，需要从最大宽度deep 逐个读取
		level, ok := levels[i]
		if !ok {
			continue
		}
		levelVals := make([]int, 0)
		for j := 0; j <= deep; j++ {
			vals, ok := level[j]
			if !ok {
				continue
			}
			// 出现在同一个位置的值按照大小追加
			sort.Ints(vals)
			levelVals = append(levelVals, vals...)
		}
		res = append(res, levelVals)
	}
	return res
}

type Node struct {
	Col int
	Row int
	Val int
}

var nodes []Node

func dfs987(root *TreeNode, row, col int) {
	if root == nil {
		return
	}
	v := Node{
		Col: col,
		Row: row,
		Val: root.Val,
	}
	nodes = append(nodes, v)
	dfs987(root.Left, row+1, col-1)
	dfs987(root.Right, row+1, col+1)
}

func verticalTraversal2(root *TreeNode) [][]int {
	nodes = []Node{}
	dfs987(root, 0, 0)
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Col != nodes[j].Col {
			return nodes[i].Col < nodes[j].Col
		}
		if nodes[i].Row != nodes[j].Row {
			return nodes[i].Row < nodes[j].Row
		}
		return nodes[i].Val < nodes[j].Val
	})

	var res [][]int
	var ans []int
	var col = nodes[0].Col

	for _, n := range nodes {
		if n.Col != col {
			col = n.Col
			res = append(res, ans)
			ans = []int{}
		}
		ans = append(ans, n.Val)
	}
	res = append(res, ans)
	return res
}
