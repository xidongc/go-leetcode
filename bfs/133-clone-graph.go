package main

import (
	"container/list"
)


// Definition for a Node.
type Node struct {
	Val int
	Neighbors []*Node
}

// bfs to traverse graph, use hashmap to store nodes relationship
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	nodesRef := make(map[*Node]*Node)
	queue := list.New()
	queue.PushBack(node)
	newHead := copyNode(node)
	nodesRef[node] = newHead

	for queue.Len() > 0 {
		ele := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		for _, neighbor := range ele.Neighbors {
			if neighbor != nil {
				if _, exist := nodesRef[neighbor]; !exist{
					queue.PushBack(neighbor)
					nodesRef[neighbor] = copyNode(neighbor)
				}
			}
		}
	}

	for original, clone := range nodesRef {
		var newNeighbors []*Node
		for _, neighbor := range original.Neighbors {
			newNeighbors = append(newNeighbors, nodesRef[neighbor])
		}

		clone.Neighbors = newNeighbors
	}


	return newHead
}

// helper func for creating new node without neighbors
func copyNode(node *Node) (newNode *Node) {
	newNode = &Node{
		Val: node.Val,
		Neighbors: []*Node{},
	}
	newNode.Val = node.Val
	return
}
