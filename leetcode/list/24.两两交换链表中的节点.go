package leetcode
// [24] 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}
	cur := head
	pre := dummyHead
	for cur != nil && cur.Next != nil {
		tmp := cur.Next.Next
		cur.Next.Next = cur
		pre.Next = cur.Next
		cur.Next = tmp
		pre = cur
		cur = tmp
	}
	return dummyHead.Next
}
