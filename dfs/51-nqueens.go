package main

// 51, 52 NQueues: use permutation to make sure | and --
// canAttach func make sure no \ and / attack
// solutions is a subnet of permutation
func solveNQueens(n int) [][]string {
	sig := make([]bool, n, n)
	result := make([][]int, 0)
	graphResult := make([][]string, 0)
	NQueuesHelper([]int{}, &result, n, sig)
	visualizeResults(&result, &graphResult, n)
	return graphResult
}

// dfs
func NQueuesHelper(solution []int, solutions *[][]int, n int, sig []bool) {
	if len(solution) == n {
		tmp := make([]int, len(solution), len(solution))
		copy(tmp, solution)
		*solutions = append(*solutions, tmp)
	}

	for i := 0 ; i < n ; i ++ {
		if sig[i] == true || canAttack(solution, i) {
			continue
		}
		solution = append(solution, i)
		sig[i] = true
		NQueuesHelper(solution, solutions, n, sig)
		solution = solution[:len(solution) - 1]
		sig[i] = false
	}
}

func canAttack(solution []int, val int) bool {
	for i := 0 ; i < len(solution); i ++ {
		if val - len(solution) + i == solution[i] ||
			val + len(solution) - i == solution[i] {
			return true
		}
	}
	return false
}

func visualizeResult(result *[]int, graphResult *[]string, n int) {
	for _, line := range *result {
		lineResult := make([]byte, n, n)
		for i := 0; i < n; i++ {
			lineResult[i] = '.'
		}
		lineResult[line] = 'Q'
		*graphResult = append(*graphResult, string(lineResult))
	}
}

func visualizeResults(result *[][]int, graphResult *[][]string, n int) {
	for _, r := range *result {
		gResult := make([]string, 0)
		visualizeResult(&r, &gResult, n)
		*graphResult = append(*graphResult, gResult)
	}
}
