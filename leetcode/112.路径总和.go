package leetcode

// [112] 路径总和

func hasPathSumR(root *TreeNode, targetSum int) bool {
	var tranversal  func(n *TreeNode, sum int) bool
	tranversal = func(n *TreeNode, sum int) bool {
		if n == nil {
			return false
		}
		if sum + n.Val == targetSum && n.Left == nil && n.Right == nil{
			return true
		}
		if n.Left != nil {
			l := tranversal(n.Left, sum + n.Val)
			if l  {
				return true
			}
		}
		if n.Right != nil {
			r := tranversal(n.Right, sum + n.Val)
			if r  {
				return true
			}
		}
		return false
	}
	return tranversal(root, 0)
}


func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	q := &Stack{}
	sum := &Stack{}
	q.Push(root)
	sum.Push(root.Val)
	for !q.IsEmpty() {
		n := q.Pop().(*TreeNode)
		s := sum.Pop().(int)
		if n.Left == nil && n.Right == nil {
			if s == targetSum {
				return true
			}
		}
		if n.Left != nil {
			q.Push(n.Left)
			sum.Push(s + n.Left.Val)
		}
		if n.Right != nil {
			q.Push(n.Right)
			sum.Push(s + n.Right.Val)
		}
	}
	return false

}