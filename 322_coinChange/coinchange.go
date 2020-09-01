package main

import (
	"container/list"
	"math"
	"sort"
)

// 解法一：暴力递归
// note：子问题总数为递归树节点个数，这个比较难看出来，总之是指数级别的，是 O(n^k)。
// 每个子问题中含有一个 for 循环，复杂度为 O(k)。所以总时间复杂度为 O(k * n^k)，指数级别。
func coinChange0(coins []int, amount int) int {
	// 当递归返回-1时，表明无法凑出零钱
	if amount < 0 {
		return -1
	}
	// 当递归时，amount金额为0时，返回0表示不需要再凑零钱了
	if amount == 0 {
		return 0
	}
	// 求最小值，所以初始化为MaxInt64
	result := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		subResult := coinChange0(coins, amount-coins[i])
		// 子问题无解，跳过
		if subResult == -1 {
			continue
		}
		result = min(result, 1+subResult)
	}
	if result == math.MaxInt64 {
		return -1
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 解法二：带备忘录的递归（自顶向下的解法）。时间复杂度为O(kn)
func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	return coinChange1Core(coins, amount, dp)
}

// amount即为不断求解的剩下的金额
func coinChange1Core(coins []int, amount int, dp []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if dp[amount] != 0 {
		return dp[amount]
	}
	result := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		subResult := coinChange1Core(coins, amount-coins[i], dp)
		if subResult >= 0 && subResult < result {
			result = 1 + subResult
		}
	}
	if result == math.MaxInt64 {
		return -1
	}
	dp[amount] = result
	return result
}

// 解法三：带备忘录的迭代（自底向上的解法）。时间复杂度为O(kn)
// note：dp数组的定义：当目标金额为i时，至少需要dp[i]枚硬币凑出。
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	// 因为凑成amount金额的硬币数最多只可能等于amount（全用1元面值的硬币），
	// 所以初始化为amount+1就相当于初始化为正无穷，便于后续取最小值。
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coins[j]])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 解法四：bfs
// note：https://leetcode-cn.com/problems/coin-change/solution/dong-tai-gui-hua-shi-yong-wan-quan-bei-bao-wen-ti-/
func coinChange3(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	queue := list.New()
	queue.PushBack(amount)
	visited := make([]bool, amount+1)
	visited[amount] = true

	sort.Ints(coins)
	result := 1

	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			head := queue.Front()
			queue.Remove(head)
			for j := 0; j < len(coins); j++ {
				leftAmount := head.Value.(int) - coins[j]
				if leftAmount == 0 {
					return result
				}
				if leftAmount < 0 {
					break
				}
				if !visited[leftAmount] {
					queue.PushBack(leftAmount)
					visited[leftAmount] = true
				}
			}
		}
		result++
	}
	return -1
}

// 解法五：dfs
// note：https://leetcode-cn.com/problems/coin-change/solution/322-by-ikaruga/
func coinChange4(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	// 金额面值由大到小排序，优先使用最大面值来凑
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	result := math.MaxInt64
	coinChangeDFS(coins, amount, 0, 0, &result)
	if result == math.MaxInt64 {
		return -1
	}
	return result
}
func coinChangeDFS(coins []int, amount int, coinIdx int, coinCnt int, result *int) {
	if amount == 0 {
		*result = min(*result, coinCnt)
		return
	}
	if coinIdx == len(coins) {
		return
	}
	// 剪枝条件：coinCnt + cnt >= *result
	// coinCnt + cnt >= *result表示加上需要更多的硬币数(cnt)超过当前的硬币数，明显不符合需要最小硬币数的条件
	for cnt := amount / coins[coinIdx]; cnt >= 0 && coinCnt+cnt < *result; cnt-- {
		leftAmount := amount - cnt*coins[coinIdx]
		// 避免影响coinCnt
		newCoinCnt := coinCnt + cnt
		coinChangeDFS(coins, leftAmount, coinIdx+1, newCoinCnt, result)
	}
}
