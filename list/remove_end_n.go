package list
// 19
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	vHead := &ListNode{
		Next: head,
	}
	fast, slow := vHead, vHead
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
        slow = slow.Next
		fast = fast.Next
	}
    slow.Next = slow.Next.Next
	return vHead.Next
}


func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	vHead := &ListNode{
		Next: head,
	}
	slow := vHead
	fast := vHead
	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			return head
		}
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return vHead.Next
}