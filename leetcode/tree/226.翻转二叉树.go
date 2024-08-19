package leetcode
// [226] 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	q := &Queue{}
	q.Push(root)
	for !q.IsEmpty() {
		n := q.Pop()
		n.Left, n.Right = n.Right, n.Left
		if n.Left != nil {
			q.Push(n.Left)
		}
		if n.Right != nil {
			q.Push(n.Right)
		}
	}
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}