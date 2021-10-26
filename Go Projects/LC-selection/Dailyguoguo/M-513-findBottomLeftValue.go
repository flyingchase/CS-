package Dailyguoguo

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	queue := make([]*TreeNode, 0)
	cur := root
	queue = append(queue, cur)
	res := make([][]int, 0)
	for len(queue) > 0 {
		lists := make([]int, 0)
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			lists = append(lists, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, lists)
	}
	return res[len(res)-1][0]
}
