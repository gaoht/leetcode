package leetcode
// [617] 合并二叉树

func mergeTreesR(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	n := &TreeNode{}
	if root1 != nil && root2 != nil {
		n.Val = root1.Val + root2.Val
		n.Left = mergeTreesR(root1.Left, root2.Left)
		n.Right = mergeTreesR(root1.Right, root2.Right)
	} else if root1  != nil {
		n.Val = root1.Val
		n.Left = mergeTreesR(root1.Left, nil)
		n.Right = mergeTreesR(root1.Right, nil)
	} else {
		n.Val = root2.Val
		n.Left = mergeTreesR(root2.Left, nil)
		n.Right = mergeTreesR(root2.Right, nil)
	}
	return n
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	q1 := &Queue{}
	q2 := &Queue{}
	q1.Push(root1)
	q1.Push(root2)
	var root *TreeNode
	for !q1.IsEmpty() {
		l1 := q1.Len()
		for i := 0; i < l1; i += 2 {
			r1 := q1.Pop()
			r2 := q1.Pop()
			n := q2.First()
			if i / 2 % 2 == 1 {
				n = q2.Pop()
			}
			if r1 == nil && r2 == nil {
				continue
			}
			var val int
			if r1 != nil && r2 != nil{
				val = r1.Val + r2.Val
				q1.Push(r1.Left)
				q1.Push(r2.Left)
				q1.Push(r1.Right)
				q1.Push(r2.Right)
			} else if r1 != nil && r2 == nil {
				val = r1.Val
				q1.Push(r1.Left)
				q1.Push(nil)
				q1.Push(r1.Right)
				q1.Push(nil)
			} else {
				val = r2.Val
				q1.Push(nil)
				q1.Push(r2.Left)
				q1.Push(nil)
				q1.Push(r2.Right)
			}
			if n == nil {
				root = &TreeNode{
					Val: val,
				}
				q2.Push(root)
				continue
			}
			if i / 2 % 2 == 0 {
				n.Left = &TreeNode{
					Val: val,
				}
				q2.Push(n.Left)
			} else {
				n.Right = &TreeNode{
					Val: val,
				}
				q2.Push(n.Right)
			}
		}	
	}
	return root
}

