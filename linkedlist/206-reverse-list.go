package linkedlist

// 206 reverse list I
func reverseList(head *ListNode) *ListNode {

	prevHead :=(*ListNode)(nil)
	for head != nil {
		tmp := head.Next
		head.Next = prevHead
		prevHead = head
		head = tmp
	}
	return prevHead
}

// 92 reverse list II
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	head = dummyNode

	for i := 1; i < m; i ++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	prevM := head
	head = head.Next
	M := head
	prev := (*ListNode)(nil)
	for i := m; i <= n; i ++ {
		if head == nil {
			return nil
		}
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	prevM.Next = prev
	M.Next = head

	return dummyNode.Next
}

// 25 reverse nodes in k-group
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	start := dummyNode

	for head != nil {
		sentinel := head
		for i := 0; i < k; i ++ {
			if sentinel == nil {
				start.Next = head
				return dummyNode.Next
			}
			sentinel = sentinel.Next
		}

		first := head
		prev := (*ListNode)(nil)
		for i := 0; i < k; i ++ {
			tmp := head.Next
			head.Next = prev
			prev = head
			head = tmp
		}
		start.Next = prev
		start = first
	}
	return dummyNode.Next
}
