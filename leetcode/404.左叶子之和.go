package leetcode
// [404] 左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	return tranversal(root, false)
}

func tranversal(n *TreeNode, left bool) int{
	if n == nil {
		return 0
	}
	if left && n.Left == nil && n.Right == nil {
		return n.Val
	}
	s := 0
	if n.Left != nil {
		s += tranversal(n.Left, true)
	}
	if n.Right != nil {
		s += tranversal(n.Right, false)
	}
	return s
}


