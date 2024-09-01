package leetcode

// [106] 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(postorder) != len(inorder) {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}

	r := postorder[len(postorder) - 1]
	root := &TreeNode{
		Val: r,
	}
	i := 0
	// 查找分隔中序
	for ; i < len(inorder); i++{
		if inorder[i] == r {
			break
		}
	}
	inorderLeft := inorder[:i]
	var inorderRight []int
	if i + 1 <= len(inorder) {
		inorderRight = inorder[i+1:]
	}
	
	root.Left = buildTree(inorderLeft, postorder[:len(inorderLeft)])
	root.Right = buildTree(inorderRight, postorder[len(inorderLeft) : len(postorder) - 1])
	return root
}
