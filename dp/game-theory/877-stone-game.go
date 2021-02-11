package game_theory

import (
	"github.com/xidongc/go-leetcode/utils"
	"math"
	"strconv"
)

/*
	877. stone game
	Alex and Lee play a game with piles of stones.  There are an even number of piles arranged in a row,
	and each pile has a positive integer number of stones piles[i].
	The objective of the game is to end with the most stones.  The total number of stones is odd, so there are no ties.
	Alex and Lee take turns, with Alex starting first.  Each turn, a player takes the entire pile of stones
	from either the beginning or the end of the row.  This continues until there are no more piles left,
	at which point the person with the most stones wins.
	Assuming Alex and Lee play optimally, return True if and only if Alex wins the game.

	game dp + interval dp, O(n^2) although it can be solved by simply "return true",
	suggest use dp as a common solution to solve this kinds of problem
	win condition: dp[0][len(piles)-1] > half

Example 1:
Input: piles = [5,3,4,5]
Output: true
Explanation:
Alex starts first, and can only take the first 5 or the last 5.
Say he takes the first 5, so that the row becomes [3, 4, 5].
If Lee takes 3, then the board is [4, 5], and Alex takes 5 to win with 10 points.
If Lee takes the last 5, then the board is [3, 4], and Alex takes 4 to win with 9 points.
This demonstrated that taking the first 5 was a winning move for Alex, so we return true.
*/
func stoneGame(piles []int)bool {
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

/*
	1563 stone game V
	There are several stones arranged in a row, and each stone has an associated value which is an integer given in the array stoneValue.
	In each round of the game, Alice divides the row into two non-empty rows (i.e. left row and right row), then Bob calculates the
	value of each row which is the sum of the values of all the stones in this row. Bob throws away the row which has the maximum value, and Alice's score increases by the value of the remaining row. If the value of the two rows are equal, Bob lets Alice decide which row will be thrown away. The next round starts with the remaining row.
	The game ends when there is only one stone remaining. Alice's is initially zero.
	Return the maximum score that Alice can obtain.

	preSum + interval dp O(n^3):
	length == 0 dp[i][i] == 0 instead of stoneValue[i], because sep will cover condition when one subarray contains only one ele
	sep <= start+length, subarray [start, sep) [sep, start+length+1)

Example 1:
Input: stoneValue = [6,2,3,4,5,5]
Output: 18
Explanation: In the first round, Alice divides the row to [6,2,3], [4,5,5]. The left row has the value 11 and the right row has value 14. Bob throws away the right row and Alice's score is now 11.
In the second round Alice divides the row to [6], [2,3]. This time Bob throws away the left row and Alice's score becomes 16 (11 + 5).
The last round Alice has only one choice to divide the row which is [2], [3]. Bob throws away the right row and Alice's score is now 18 (16 + 2). The game ends because only one stone is remaining in the row.

Example 2:
Input: stoneValue = [7,7,7,7,7,7,7]
Output: 28
Example 3:
Input: stoneValue = [4]
Output: 0
*/
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
// typical game-theory with 2*M fan out, don't forget minVal = 0, it cost me hours to debug my solution
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

