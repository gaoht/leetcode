package leetcode

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxPathSum(root *TreeNode) int {
    sum := 0
    max := math.MinInt
    var inorder func(node *TreeNode, right bool)
    inorder = func(node *TreeNode, right bool) {
        if node == nil {
            return
        }
        if node.Left != nil {
            inorder(node.Left, false)
        }
        fmt.Println("val:", node.Val, "sum:", sum, "max:", max)
        //当前节点
        sum += node.Val
        if max < sum {
            max = sum
        }
        if sum < 0 {
            sum = 0
        }
        if node.Right != nil {
            inorder(node.Right, true)
        } else if right {
            sum = 0
        }   
    }
    inorder(root, false)
    return max
}