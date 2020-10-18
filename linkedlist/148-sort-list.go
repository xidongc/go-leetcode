package linkedlist

// 148 sort list using merge sort
func sortListMS(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	list1 := sortListMS(slow.Next)
	slow.Next = nil
	list2 := sortListMS(head)
	return mergeTwoLists(list1, list2)
}

// 148 sort list using quick sort
func sortListQS(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	less, equal, larger := &ListNode{}, &ListNode{}, &ListNode{}
	dummyLess, dummyEqual, dummyLarger := less, equal, larger

	cmp := head.Val
	for head != nil {
		if head.Val == cmp {
			equal.Next = head
			equal = equal.Next
		} else if head.Val < cmp {
			less.Next = head
			less = less.Next
		} else {
			larger.Next = head
			larger = larger.Next
		}
		head = head.Next
	}
	less.Next = nil
	larger.Next = nil

	left := sortListQS(dummyLess.Next)
	right := sortListQS(dummyLarger.Next)

	if left != nil {
		lastTail(left).Next = dummyEqual.Next
	} else {
		left = dummyEqual.Next
	}
	equal.Next = right
	return left
}

func lastTail(node *ListNode) *ListNode {
	for node != nil && node.Next != nil {
		node = node.Next
	}
	return node
}