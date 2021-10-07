package Dailyguoguo

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {

	var bi BSTIterator
	bi.push(root)
	return bi
}

func (this *BSTIterator) Next() int {

	tmp := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.push(tmp.Right)
	return tmp.Val
}

func (this *BSTIterator) HasNext() bool {

	if len(this.stack) == 0 {
		return false
	}
	return true
}

func (bi *BSTIterator) push(root *TreeNode) {
	for root != nil {
		bi.stack = append(bi.stack, root)
		root = root.Left
	}
}
