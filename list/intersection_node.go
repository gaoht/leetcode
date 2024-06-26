package list
// 160 GetIntersectionNode
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	// nil == nil 时退出
	for p1 != p2{
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

func GetIntersectionNode2(headA, headB *ListNode) *ListNode {
	lenA := listLength(headA)
	lenB := listLength(headB)
	skip := lenA - lenB
	tHeadA := headA
	tHeadB := headB
	if lenB > lenA {
		tHeadA = headB
		tHeadB = headA
		skip = lenB - lenA
	}
	i := 0
	for i < skip {
		tHeadA = tHeadA.Next
		i++
	}
	for tHeadA != nil{
		if tHeadA == tHeadB {
			return tHeadA
		} else {
			tHeadA = tHeadA.Next
			tHeadB = tHeadB.Next
		}
	}
	return nil

}

func listLength(head * ListNode) int{
	cur := head
	len := 0
	for cur != nil {
		len ++
		cur = cur.Next
	}
	return len
}