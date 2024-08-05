package leetcode

import (
	"fmt"
	"strings"
)

// ListNode 链表
type ListNode struct {
    Val int
    Next *ListNode
}

// PrintList 打印
func printList(head *ListNode) {
	x := strings.Builder{}
	for cur := head; cur != nil;  cur = cur.Next {
		x.WriteString(fmt.Sprintf("%d ", cur.Val))
	}
	fmt.Println(x.String())
}

func buildList(nums []int) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for  _, v := range nums {
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}
	return dummyHead.Next
}
 
func listToArray(head *ListNode) []int{
	ret := make([]int, 0)
	cur := head
	for cur != nil {
		ret = append(ret, cur.Val)
		cur = cur.Next
	}
	return ret
}