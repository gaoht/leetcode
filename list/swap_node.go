 package list

 func SwapPairs(head *ListNode) *ListNode {
	vHead := &ListNode{
		Next: head,
	}
	cur := vHead
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next.Next.Next
		tmp2 := cur.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp2 
		cur.Next.Next.Next = tmp1
		cur = cur.Next.Next
	}
	return vHead.Next
 }