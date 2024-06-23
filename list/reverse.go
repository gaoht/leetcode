package list

// 206
func ReverseList(head *ListNode) *ListNode {
	var pre  *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur	
		cur = tmp
	}
	return pre
}