package leetcode
// [105] 从前序与中序遍历序列构造二叉树

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder)  {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val: preorder[0],
		}
	}
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == preorder[0] {
			break
		}
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	inLeft := inorder[:index]
	var inRight []int
	if index + 1 <= len(inorder) {
		inRight = inorder[index + 1 :]
	}
	root.Left = buildTree2(preorder[1:1+len(inLeft)], inLeft)
	root.Right = buildTree2(preorder[1+len(inLeft) : ], inRight)
	return root
}

