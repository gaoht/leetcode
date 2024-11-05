package leetcode

import "fmt"

// [114] 二叉树展开为链表
func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    pre := &TreeNode {}
    var traversal func(node *TreeNode)
    traversal = func(node *TreeNode) {
        if node == nil {
            return
        }
		
        l, r := node.Left, node.Right
		pre.Right = node
		if l == nil && r == nil {
			pre = node
		}
        if l != nil {
            pre.Left = nil
            traversal(l)
			pre = l
        } 
        if r != nil {
			pre.Left = nil
            traversal(r)
			pre = r
        }
		
    }
    traversal(root)
}



func preOrder(r *TreeNode){
	if r == nil {
		return
	}
	fmt.Printf(" %d", r.Val)
	preOrder(r.Left)
	preOrder(r.Right)
}