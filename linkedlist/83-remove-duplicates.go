package linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

// 83 remove duplicates I
func deleteDuplicates(head *ListNode) *ListNode {
	start := head
	for head != nil && head.Next != nil {
		if head.Next.Val == head.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return start
}

// 82 remove duplicates II
func deleteDuplicatesII(head *ListNode) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	head = dummyNode
	for head != nil && head.Next != nil {
		current := head.Next
		if current.Next != nil && current.Next.Val == current.Val {
			for current.Next != nil && current.Next.Val == current.Val {
				current = current.Next
			}
			head.Next = current.Next
		} else {
			head = head.Next
		}}
	return dummyNode.Next
}
