package utils

import (
	"container/list"
	"fmt"
	"math"
)

const (
	NilTreeNode =  math.MinInt64  // nil tree node
)

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// Create TreeNode with val
func CreateTreeNode(val int) *TreeNode {
	return &TreeNode {
		Val: val,
		Left: nil,
		Right: nil,
	}
}

// Create tree from given array, return root
func CreateTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	queue := list.New()
	root := CreateTreeNode(arr[0])
	pos := 1
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if pos < len(arr) {
			leftNode := (*TreeNode)(nil)
			if arr[pos] != NilTreeNode {
				leftNode = CreateTreeNode(arr[pos])
				queue.PushBack(leftNode)
			}
			node.Left = leftNode
			pos += 1
		}
		if pos < len(arr) {
			rightNode := (*TreeNode)(nil)
			if arr[pos] != NilTreeNode {
				rightNode = CreateTreeNode(arr[pos])
				queue.PushBack(rightNode)
			}
			node.Right = rightNode
			pos += 1
		}
	}
	return root
}

// In order traverse used for debugging
func InOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraverse(root.Left)
	fmt.Println(root.Val)
	InOrderTraverse(root.Right)
}
