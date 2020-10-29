package main

// 572 subtree of another tree
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSubtreeHelper(s, t) {
		return true
	} else if s == nil {
		return false
	}
	leftOk := isSubtree(s.Left, t)
	rightOk := isSubtree(s.Right, t)
	return rightOk || leftOk
}

func isSubtreeHelper(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isSubtreeHelper(s.Left, t.Left) && isSubtreeHelper(s.Right, t.Right)
}
