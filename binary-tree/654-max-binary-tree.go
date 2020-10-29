package main

// 654 max binary tree
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := getMaxIndex(nums)
	root := &TreeNode{
		Val: nums[i],
		Left: constructMaximumBinaryTree(nums[0:i]),
		Right: constructMaximumBinaryTree(nums[i+1:]),
	}
	return root
}

// get max index
func getMaxIndex(nums []int) int {
	pos := 0
	for i := 1; i < len(nums); i ++ {
		if nums[i] > nums[pos] {
			pos = i
		}
	}
	return pos
}

//

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if val > root.Val {
		tmp := root.Val
		root.Val = val
		child := getMaxChild(root)
		if child == nil {
			root.Left = &TreeNode{
				val,
				nil,
				nil,
			}
		} else {
			insertIntoMaxTree(child, tmp)
		}
		return root
	} else if val < root.Val {
		child := getMaxChild(root)
		if child == nil {
			root.Left = &TreeNode{
				val,
				nil,
				nil,
			}
		} else {
			insertIntoMaxTree(child, tmp)
		}
		return root
	}
	return root
}

// get max child
func getMaxChild(root *TreeNode) *TreeNode {
	if root.Left != nil && root.Right != nil {
		if root.Left.Val > root.Right.Val {
			return root.Left
		}
		return root.Right
	} else if root.Left != nil {
		return root.Left
	} else if root.Right != nil {
		return root.Right
	}
	return nil
}

