package leetcode

// [513] 找树左下角
func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	q := &Queue{}
	q.Push(root)
	for !q.IsEmpty() {
		l := q.Len()
		var left *int
		bottom := true
		for i := 0; i < l; i++{
			n := q.Pop()
			if left == nil {
				left = &n.Val
			}
			if n.Left != nil {
				bottom = false
				q.Push(n.Left)
			}
			if n.Right != nil {
				bottom = false
				q.Push(n.Right)
			}
		}
		if bottom {
			return *left
		}
	}
	return -1
}


func findBottomLeftValueR(root *TreeNode) int{
	maxLevel := 0
	result := 0
	if root == nil {
		return result
	}
	var tranversal func(n *TreeNode, depth int)
	tranversal = func(n *TreeNode, depth int)  {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			depth += 1
			if maxLevel < depth {
				maxLevel = depth
				result = n.Val
			}
		}
		if n.Left != nil {
			tranversal(n.Left, depth + 1)
		}
		if n.Right != nil {
			tranversal(n.Right, depth + 1)
		}
	}
	tranversal(root, 0)
	return result
}

