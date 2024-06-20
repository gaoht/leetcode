package list

type ListNode struct{
	Val int
	Next *ListNode
}


// RemoveElements 203
func RemoveElements(head *ListNode, target int) *ListNode {
	if head == nil {
		return nil
	}
	vHead := &ListNode{
		Val: -1,
		Next: head,
	}
	cur := vHead
	for cur.Next != nil {
		if cur.Next.Val == target {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return vHead.Next
}