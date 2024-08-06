package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}


func newList(nums []int) *ListNode {
	dummyHead := &ListNode{
	}
	cur := dummyHead
	for _, v := range nums {
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}
	return dummyHead.Next
}

func printList(h *ListNode) {
	if h == nil {
		return
	}
	cur := h
	sb := strings.Builder{}
	for cur != nil {
		sb.WriteString(strconv.Itoa(cur.Val))
		if cur.Next != nil {
			sb.WriteByte(' ')
		}
	}
	fmt.Println(sb.String())
}

func listToArray(h *ListNode) []int{
	n := make([]int, 0)
	for cur := h; cur != nil; cur = cur.Next {
		n = append(n, cur.Val)
	}
	return n
}