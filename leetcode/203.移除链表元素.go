package leetcode

// [203] 移除链表元素

// [1,2,6,3,4,5,6]
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{
		Next: head,
	}
	cur := dummyHead
	for cur != nil && cur.Next != nil{
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			// 如果没有else会多次跳过
			// 比如 7,7,7,7		
			cur = cur.Next
		}	
	}
	return dummyHead.Next
}
