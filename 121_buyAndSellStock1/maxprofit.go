package main

// 暴力递归，时间复杂度为O(n^2)
func maxProfit0(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxP := 0
	for buy := 0; buy < len(prices); buy++ {
		for sell := buy+1; sell < len(prices); sell++ {
			maxP = max(maxP, prices[sell]-prices[buy])
		}
	}
	return maxP
}

// 动态规划，时间复杂度为O(n)
// 定义dp[i]为在前i天的最大收益
// 前i天的最大收益 = max(前i-1天的最大收益, 第i天的最大收益)
// 状态转移方程：dp[i] = max(dp[i-1], prices[i]-minPrice)
func maxProfit1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([]int, len(prices))
	minPrice := prices[0]
	for sell := 1; sell < len(prices); sell++ {
		// 找出当前最小的价格
		minPrice = min(minPrice, prices[sell])
		// 找出当前最大的利润
		dp[sell] = max(dp[sell-1], prices[sell]-minPrice)
	}
	// dp[i]总是取截止到目前为止的最大收益值，所以最后一位一定是最大的
	return dp[len(prices)-1]
}

// TODO: 二维dp
// 状态转移方程
func maxProfit2(prices []int) int {
	return 0
}

// 动态规划压缩版，时间复杂度为O(n)，站在sell的角度来解题
func maxProfit3(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxP := 0
	minPrice := prices[0]
	for sell := 1; sell < len(prices); sell++ {
		// 找出当前最小的价格
		minPrice = min(minPrice, prices[sell])
		// 找出当前最大的利润
		maxP = max(maxP, prices[sell]-minPrice)
	}
	return maxP
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}