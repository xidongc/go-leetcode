package main

import (
	"github.com/xidongc/go-leetcode/utils"
	"math"
	"strconv"
)

// 123 stock prices III O(n)
// first separate into 2 half, left and right subarray represents max profits for each. left[i] represents:
// from left take i ele, the maximum profit I can get. right[i] represents, if left already take i ele, max
// profit I can get from right separation.
// my ans: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/908954/Golang-O(n)-Solution-(NO-DP-needed)
func maxProfitTwoTransaction(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	left := make([]int, len(prices)+1, len(prices)+1)
	leftMin := prices[0]

	for i := 2; i <= len(prices); i ++ {
		left[i] = utils.Max(left[i-1], prices[i-1] - leftMin)
		leftMin = utils.Min(prices[i-1], leftMin)
	}

	right := make([]int, len(prices)+1, len(prices)+1)
	rightMax := prices[len(prices)-1]

	for i := len(prices); i >=2 ; i-- {
		right[i-2] = utils.Max(right[i-1], rightMax - prices[i-2])
		rightMax = utils.Max(prices[i-2], rightMax)
	}

	maxValue := math.MinInt64
	for i := 0; i <= len(prices); i ++ {
		maxValue = utils.Max(maxValue, right[i]+left[i])
	}
	return maxValue
}

// 122 stock prices II
// transaction as many time as possible to make max profits
// one pass O(n)
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProfits := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i ++ {
		if prices[i] - minPrice > 0 {
			maxProfits += prices[i] - minPrice
		}
		minPrice = prices[i]
	}
	return maxProfits
}

// 121 stock prices I
// transaction one time, one pass O(n)
func maxProfitOneTransaction(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i ++ {
		maxProfit = utils.Max(maxProfit, prices[i]-minPrice)
		minPrice = utils.Min(minPrice, prices[i])
	}
	return maxProfit
}

// 188 stock price IV
// transaction k times using memorized divide conquer, LTE O((n-2k)*n*log(k))
func maxProfitKTransaction(k int, prices []int) int {
	return maxProfitKTransactionMemo(k, prices, map[string]int{})
}

func maxProfitKTransactionMemo(k int, prices []int, memo map[string]int) int {
	code := hashCode(prices, k)
	if val, exist := memo[code]; exist {
		return val
	}

	if k == 2 {
		result := maxProfitTwoTransaction(prices)
		memo[code] = result
		return result
	} else if k == 1 {
		result := maxProfitOneTransaction(prices)
		memo[code] = result
		return result
	} else if k < 1 || len(prices) < 2 {
		return 0
	}
	sep := k/2
	maxProfit := 0
	for i := 0; i <= len(prices); i ++ {
		if i < sep * 2 || len(prices) - i < (k-sep)*2 {
			continue
		}
		maxProfit = utils.Max(maxProfit, maxProfitKTransactionMemo(sep, prices[0:i], memo) + maxProfitKTransactionMemo(k-sep, prices[i:], memo))
	}
	memo[code] = maxProfit
	return maxProfit
}

func hashCode(input []int, k int) string {
	result := ""
	for _, i := range input {
		result += strconv.Itoa(i)
		result += "s"
	}
	result += "x"
	result += strconv.Itoa(k)
	return result
}
