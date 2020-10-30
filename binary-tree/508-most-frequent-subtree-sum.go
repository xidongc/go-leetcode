package main

// 508 most frequent subtree sum: using divide conquer
func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	dict, _ := findFrequentHelper(root)
	results := make([]int, 0)
	cmp := 1
	for key, value := range dict {
		if value > cmp {
			results = []int{key}
			cmp = value
		} else if value == cmp {
			results = append(results, key)
		}
	}
	return results
}

func findFrequentHelper(root *TreeNode) (map[int]int, int) {
	if root == nil {
		return map[int]int{}, 0
	}

	leftResults,  left  := findFrequentHelper(root.Left)
	rightResults, right := findFrequentHelper(root.Right)

	results := make(map[int]int, 0)
	for key, value := range leftResults {
		results[key] = value
	}

	for key, value := range rightResults {
		if _, exist := results[key]; exist {
			results[key] += value
		} else {
			results[key] = value
		}
	}

	current := root.Val + left + right
	if val, exist := results[current]; exist {
		results[current] = val + 1
	} else {
		results[current] = 1
	}

	return results, current
}
