package leetcode
// [141] 环形链表


func hasCycle(head *ListNode) bool {
    fast, slow := head, head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

