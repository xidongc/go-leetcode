package game_theory

import (
	"github.com/xidongc/go-leetcode/utils"
	"math"
	"strconv"
)

// 877 stone game
// game dp + interval dp, O(n^2) although it can be solved by simply "return true",
// suggest use dp as a common solution to solve this kinds of problem
// win condition: dp[0][len(piles)-1] > half
func stoneGame(piles []int) bool {
	dp := make([][]int, 0)
	sum := 0
	for i := 0; i < len(piles); i ++ {
		sum += piles[i]
		dp = append(dp, make([]int, len(piles), len(piles)))
	}

	for length := 0; length < len(piles); length ++ {
		for start := 0; start < len(piles) - length; start ++ {
			if length == 0 {
				dp[start][start] = piles[start]
			} else if length == 1 {
				dp[start][start+1] = utils.Max(piles[start], piles[start+1])
			} else {
				dp[start][start+length] = utils.Max(piles[start]+utils.Min(dp[start+1][start+length-1],
					dp[start+2][start+length]), piles[start+length]+utils.Min(dp[start+1][start+length-1],
					dp[start][start+length-2]))
			}
		}
	}
	return dp[0][len(piles)-1] > (sum / 2)
}

// 1563 stone game V
// preSum + interval dp O(n^3), some small points might neglect:
// length == 0 dp[i][i] == 0 instead of stoneValue[i], because sep
// will cover condition when one subarray contains only one ele
// also sep <= start+length, subarray [start, sep) [sep, start+length+1)
func stoneGameV(stoneValue []int) int {
	preSum := make([]int, len(stoneValue)+1, len(stoneValue)+1)
	for i := 0; i < len(stoneValue); i ++ {
		preSum[i+1] = preSum[i] + stoneValue[i]
	}
	dp := make([][]int, 0)
	for i := 0; i < len(stoneValue); i ++ {
		dp = append(dp, make([]int, len(stoneValue), len(stoneValue)))
	}
	for length := 0; length < len(stoneValue); length ++ {
		for start := 0; start < len(stoneValue) - length; start ++ {
			if length == 0 {
				dp[start][start+length] = 0
				continue
			} else if length == 1 {
				dp[start][start+length] = utils.Min(stoneValue[start], stoneValue[start+length])
			}
			for sep := start+1; sep <= start+length; sep ++ {
				left := preSum[sep] - preSum[start]
				right := preSum[start+length+1] - preSum[sep]
				if left < right {
					dp[start][start+length] = utils.Max(dp[start][start+length], dp[start][sep-1] + left)
				} else if left > right {
					dp[start][start+length] = utils.Max(dp[start][start+length], dp[sep][start+length] + right)
				} else {
					dp[start][start+length] = utils.Max(dp[start][sep-1] + left, dp[sep][start+length] + right)
				}
			}
		}
	}
	return dp[0][len(stoneValue)-1]
}

// 1140 stone game II
// typical game theory with 2*M fan out, don't forget minVal = 0, it cost me hours to debug my solution
// use start and M to memorize, previous use array slice, which lead to TLE
func stoneGameII(piles []int) int {
	return stoneGameIIHelper(piles, 0, 1, map[string]int{})
}

func stoneGameIIHelper(piles []int, start int, M int, hash map[string]int) int {
	if val, exist := hash[cusHash(start, M)]; exist {
		return val
	}
	if len(piles) - start <= 2*M {
		sum := 0
		for i := start; i < len(piles); i ++ {
			sum += piles[i]
		}
		return sum
	}
	value := 0
	maxVal := 0
	for i := start + 1; i <= start + 2 * M; i ++ {
		if i >= len(piles) {
			break
		}
		value += piles[i-1]
		minVal := math.MaxInt64
		fanOut := utils.Max(M, i - start)
		for j := 1; j <= 2 *fanOut; j ++ {
			if i+j >= len(piles) {
				minVal = 0
				break
			}
			minVal = utils.Min(minVal, stoneGameIIHelper(piles, i+j, utils.Max(fanOut, j), hash))
		}
		maxVal = utils.Max(maxVal, value + minVal)
	}
	hash[cusHash(start, M)] = maxVal
	return maxVal
}

func cusHash(input int, M int) string {
	output := strconv.Itoa(input)
	output += "$"
	output += strconv.Itoa(M)
	output += "@"
	return output
}

