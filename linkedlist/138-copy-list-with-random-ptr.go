package linkedlist


// Definition for a Node.
type Node struct {
	Val int
	Next *Node
	Random *Node
}


// solution with map
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dummyNode := &Node{}
	newHead := dummyNode
	current := head
	newToOld := make(map[*Node]*Node, 0)
	oldToNew := make(map[*Node]*Node, 0)

	for current != nil {
		newHead.Next = &Node{
			Val: current.Val,
			Random: current,
		}
		newToOld[newHead.Next] = current
		oldToNew[current] = newHead.Next
		current = current.Next
		newHead = newHead.Next
	}
	newHead.Next = nil

	current = dummyNode.Next
	for current != nil {
		if newToOld[current].Random == nil {
			current.Random = nil
		} else {
			current.Random = oldToNew[newToOld[current].Random]
		}
		current = current.Next
	}
	return dummyNode.Next
}

