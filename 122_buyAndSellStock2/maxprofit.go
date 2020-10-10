package main

// 贪心算法，时间复杂度为O(n)
// 可能跨越多天的买卖都化解成相邻两天的买卖
// note：(prices[1]-prices[0])+(prices[2]-prices[1])
func maxProfit0(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxP := 0
	for sell := 1; sell < len(prices); sell++ {
		maxP += max(prices[sell]-prices[sell-1], 0)
	}
	return maxP
}

// 二维dp
func maxProfit2(prices []int) int {
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
