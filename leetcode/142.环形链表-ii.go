package leetcode
// [142] 环形链表 II

func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			x := head
			y := fast
			for x != y {
				x = x.Next
				y = y.Next
			}
			return x
		}
	}
	return nil
}
// @lc code=end

