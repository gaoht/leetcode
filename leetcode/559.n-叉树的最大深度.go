package leetcode


// [559] N 叉树的最大深度

func maxNDepth(root *Node) int {
    return maxNDepthR(root, 0)
}

func maxNDepthR(n *Node, depth int) int{
	if n == nil {
		return depth
	}
	depthx := []int{}
	if len(n.Children) == 0 {
		return depth + 1
	}
	for _, c := range n.Children {
		depthx = append(depthx, maxNDepthR(c, depth + 1))
	}
	return max(depthx)
}

func max(nums []int) int{
	x := 0
	for _, n := range nums {
		if x < n {
			x = n
		}
	}
	return x
}
