package leetcode
// [104] 二叉树的最大深度

func maxDepth(root *TreeNode) int {
	 return maxDepthR(root, 0)
}

func maxDepthR(n *TreeNode, depth int) int{
	if n == nil {
		return depth
	}
	l := maxDepthR(n.Left, depth + 1)
	r := maxDepthR(n.Right, depth + 1)

	if l > r {
		return l
	}
	return r
}

