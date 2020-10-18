package linkedlist

import (
	"container/heap"
)

// 21 merge two sorted lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ptr1 := l1
	ptr2 := l2
	dummyNode := &ListNode{}
	head := dummyNode

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val <= ptr2.Val {
			head.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			head.Next = ptr2
			ptr2 = ptr2.Next
		}
		head = head.Next
	}

	current := ptr1
	if ptr2 != nil {
		current = ptr2
	}
	for current != nil {
		head.Next = current
		current = current.Next
		head = head.Next
	}
	return dummyNode.Next
}

// 23 merge k sorted lists (priority queue solution)
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func mergeKListsHeap(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	for _, node := range lists {
		if node != nil {
			pq = append(pq, node)
		}
	}
	heap.Init(&pq)
	dummyNode := &ListNode{}
	head := dummyNode
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*ListNode)
		head.Next = item
		if item.Next != nil {
			heap.Push(&pq, item.Next)
		}
		head = head.Next
	}
	return dummyNode.Next
}

// 23 merge k sorted lists (divide conquer solution)
func mergeKListsDC(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	list1 := mergeKListsDC(lists[:mid])
	list2 := mergeKListsDC(lists[mid:])
	return mergeTwoLists(list1, list2)
}

// 23 merge k sorted lists (divide conquer solution II)
func mergeKListsDCII(lists []*ListNode) *ListNode {
	result := lists
	for len(result) > 1 {
		result = mergeKListsHelper(result)
	}
	if len(result) == 1 {
		return result[0]
	}
	return nil
}

func mergeKListsHelper(lists []*ListNode) []*ListNode {
	remaining := make([]*ListNode, 0)
	ptr1 := 0
	// odd size
	if len(lists) % 2 == 1 {
		remaining = append(remaining, lists[0])
		ptr1 = 1
	}
	ptr2 := len(lists) - 1
	for ptr1 < ptr2 {
		remaining = append(remaining, mergeTwoLists(lists[ptr1], lists[ptr2]))
		ptr1 += 1
		ptr2 -= 1
	}
	return remaining
}