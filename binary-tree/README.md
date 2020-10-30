# binary tree 

two way divide & conquer, and combine them into one

```go 
func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0 
    }
    _, max := maxPathSumHelper(root)
    return max 
}

// return singlePath, maxPath
func maxPathSumHelper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, math.MinInt64
	}
	leftSingle, leftMax := maxPathSumHelper(root.Left)
	rightSingle, rightMax := maxPathSumHelper(root.Right)
	count := root.Val
	if leftSingle > 0 {
		count += leftSingle
	}
	if rightSingle > 0 {
		count += rightSingle
	}
	single := Max(root.Val, leftSingle+root.Val, rightSingle+root.Val)
	return single, Max(leftMax, rightMax, count)
}

```

