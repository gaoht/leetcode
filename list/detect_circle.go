package list
 // DetectCycle 142 
 func DetectCycle2(head *ListNode) *ListNode {
    set := make(map[int]bool)
    cur := head
    for cur != nil {
        if ok := set[cur.Val]; !ok {
            set[cur.Val] = true
            cur = cur.Next
        } else {
            return cur
        } 
    }
    return nil
}

func DetectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			index1 := fast
			index2 := head
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index2
		}
	}
	return nil
}