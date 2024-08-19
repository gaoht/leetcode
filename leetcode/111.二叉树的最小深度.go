package leetcode
// [111] 二叉树的最小深度

func minDepth(root *TreeNode) int {
	return minDepthR(root, 0)
}

func minDepthR(n *TreeNode, depth int) int {
	if n == nil {
		return depth
	}
	if n.Left == nil && n.Right == nil {
		return depth + 1
	}
	if n.Left == nil {
		return minDepthR(n.Right, depth + 1)
	}
	if n.Right == nil {
		return minDepthR(n.Left, depth + 1)
	}
	l := minDepthR(n.Left, depth + 1)
	r := minDepthR(n.Right, depth + 1)
	if l > r {
		return r
	}
	return l
}

