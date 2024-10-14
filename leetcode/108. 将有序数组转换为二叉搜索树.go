package leetcode

import "fmt"

func SortedArrayToBST(nums []int) *TreeNode {
    numsQ := [][]int{}
    nodeQ := []*TreeNode{}
    numsQ = append(numsQ, nums) 
    var root *TreeNode
    for len(numsQ) > 0 {
        l := len(numsQ)
        for i := 0; i < l; i++ {
            array := numsQ[0]
            numsQ = numsQ[1:]
			fmt.Printf("%v", numsQ)
            var node, parent *TreeNode
            if len(array) != 0 {
                mid := len(array) / 2 
                // 分割数组
                numsQ = append(numsQ, array[0:mid])
                numsQ = append(numsQ, array[mid+1 : ])
                node = &TreeNode {
                    Val: array[mid],
                }
            }
            if len(nodeQ) == 0 {
                root = node
            } else {
                parent = nodeQ[0]
                nodeQ = nodeQ[1 : ]
                if parent != nil {
                    if i % 2 == 0 {
                        parent.Left = node  
                    }else {
                        parent.Right = node
                    }
                }
            }
            nodeQ = append(nodeQ, node)
            nodeQ = append(nodeQ, node)
        }
    }
    return root
}