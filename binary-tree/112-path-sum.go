package main

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 112 path sum I: using divide conquer
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return hasPathSumHelper(root, sum)
}

func hasPathSumHelper(root *TreeNode, value int) bool {
	if root.Left == nil && root.Right == nil {
		return value == root.Val
	}
	leftMatch, rightMatch := false, false
	if root.Left != nil {
		leftMatch = hasPathSum(root.Left, value-root.Val)
	}
	if root.Right != nil {
		rightMatch = hasPathSum(root.Right, value - root.Val)
	}
	return leftMatch || rightMatch
}

// 113 path sum II: divide conquer
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	return PathSumHelper(root, sum)
}

func PathSumHelper(root *TreeNode, value int) [][]int {
	sequence := make([][]int, 0)
	if root.Left == nil && root.Right == nil {
		if value == root.Val {
			sequence = append(sequence, []int{root.Val})
			return sequence
		} else {
			return nil
		}
	}
	if root.Left != nil {
		leftSeq := PathSumHelper(root.Left, value-root.Val)
		if leftSeq != nil {
			for _, left := range leftSeq {
				left = append([]int{root.Val}, left...)
				sequence = append(sequence, left)
			}
		}

	}
	if root.Right != nil {
		rightSeq := PathSumHelper(root.Right, value-root.Val)
		if rightSeq != nil {
			for _, right := range rightSeq {
				right = append([]int{root.Val}, right...)
				sequence = append(sequence, right)
			}
		}
	}
	return sequence
}
