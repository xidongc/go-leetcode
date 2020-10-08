package main

// search in any position
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	height := len(matrix)
	width := len(matrix[0])

	start := 0
	end := height * width - 1

	for start + 1 < end {
		mid := start + (end - start) / 2
		atHeight := mid / width
		atWidth := mid % width
		if matrix[atHeight][atWidth] == target {
			return true
		} else if matrix[atHeight][atWidth] < target {
			start = mid
		} else {
			end = mid
		}
	}
	if matrix[end / width][end % width] == target {
		return true
	}
	if matrix[start/ width][start % width] == target {
		return true
	}
	return false
}

// follow up: [240] search a 2D Matrix II
// start from either top right or bottom left to remove
// either a line or a column
func searchMatrixII(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	atHeight := len(matrix) - 1
	atWidth := 0
	for atHeight >= 0 && atWidth <= len(matrix[0]) - 1 {
		if matrix[atHeight][atWidth] == target {
			return true
		} else if matrix[atHeight][atWidth] < target {
			atWidth += 1
		} else {
			atHeight -= 1
		}
	}
	return false
}
