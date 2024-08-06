package leetcode

// [160] 相交链表
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
    sLen, lLen := lenList(headA), lenList(headB)
	sHead, lHead := headA, headB
	if sLen > lLen {
		lHead, sHead = sHead, lHead
		lLen, sLen = sLen, lLen
	}
	sCur, lCur := sHead, lHead
	for i := lLen - sLen; i > 0 && lCur != nil; i-- {
		lCur = lCur.Next
	}
	for lCur != nil{
		if lCur == sCur {
			return lCur
		}
		lCur = lCur.Next
		sCur = sCur.Next
	}
	return nil
}

func lenList(head *ListNode) int {
	l := 0
	for cur := head; cur != nil; cur = cur.Next {
		l++
	}
	return l
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
    curA, curB := headA, headB
	for curA != curB {
		if curA.Next == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB.Next == nil{
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}


