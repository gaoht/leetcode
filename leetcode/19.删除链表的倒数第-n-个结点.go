package leetcode
//[19] 删除链表的倒数第 N 个结点

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	if head == nil {
		return head
	}
	slow, fast := dummyHead, dummyHead
	i := n
	for ; i > 0 && fast.Next != nil; i--{
		fast = fast.Next
	}
	if i > 0 {
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
