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

// path sum III: two dfs recursive
func pathSumIII(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSumIII(root.Left, sum) + pathSumIII(root.Right, sum) + pathSumHelper(root, sum)
}

func pathSumHelper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0
	if sum == root.Val {
		count += 1
	}
	count += pathSumHelper(root.Left, sum-root.Val)
	count += pathSumHelper(root.Right, sum-root.Val)

	return count
}

func pathSumIIII(root *TreeNode, sum int) int {
	dict := map[int]int{}
	return pathSumHelp(root, sum, []int{}, dict)
}

func pathSumHelp(root *TreeNode, target int, preSum []int, dict map[int]int) int {
	count := 0
	for _, child := range []*TreeNode{root.Left, root.Right} {
		preSum = append(preSum, preSum[len(preSum)-1] + root.Val)
		if val, exist := dict[preSum[len(preSum)-1]]; exist {
			dict[preSum[len(preSum)-1]] = val + 1
		} else {
			dict[preSum[len(preSum)-1]] = 1
		}
		if val, exist := dict[preSum[len(preSum)-1] - target]; exist {
			count += val
		}
		if child != nil {
			pathSumHelp(child, target, preSum, dict)
		}
		delete(dict, preSum[len(preSum)-1])
		preSum = preSum[:len(preSum)-1]
	}
	return count
}

