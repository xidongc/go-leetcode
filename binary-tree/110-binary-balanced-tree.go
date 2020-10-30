package main

import (
	"github.com/xidongc/go-leetcode/utils"
	"math"
)

// 110 balanced binary tree: divide conquer
func isBalanced(root *utils.TreeNode) bool {
	_, ok := BalancedHelper(root)
	return ok
}

func BalancedHelper(root *utils.TreeNode) (height int, isBalance bool) {
	if root == nil {
		return 0, true
	}
	leftHeight, leftOk := BalancedHelper(root.Left)
	rightHeight, rightOk := BalancedHelper(root.Right)
	maxHeight := utils.Max(leftHeight, rightHeight) + 1
	if leftOk && rightOk && int(math.Abs(float64(rightHeight) - float64(leftHeight))) <= 1 {
		return maxHeight, true
	}
	return maxHeight, false
}
