package leetcode
// [437] 路径总和 III
func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	var tranversal func(root *TreeNode)
	var countPathSum func(n *TreeNode, sum int)
	countPathSum = func(n *TreeNode, sum int){
		if n == nil {
			return
		}
		sum = n.Val + sum
		if sum == targetSum {
			count++
		}
		if n.Left != nil {
			countPathSum(n.Left, sum)
		}
		if n.Right != nil {
			countPathSum(n.Right, sum)
		}
	}
	tranversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		countPathSum(root, 0)
		if root.Left != nil {
			tranversal(root.Left)
		}
		if root.Right != nil {
			tranversal(root.Right)
		}
	}
	tranversal(root)
	return count
}

