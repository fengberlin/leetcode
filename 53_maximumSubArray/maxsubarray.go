package main

import (
	"math"
)

// 解法一：动态规划，时间复杂度为O(n)
// note：最主要还是dp数组定义，这道题还不能用滑动窗口算法，因为数组中的数字可以是负数。
// dp数组定义为以nums[i]为结尾的「最大子数组和」为dp[i]。
// 状态转移方程：dp[i] = max(nums[i], nums[i]+dp[i-1])
func maxSubArray0(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	// base case
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}
	ret := nums[0]
	for i := 0; i < len(dp); i++ {
		ret = max(ret, dp[i])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法二：动态规划解法的状态压缩版
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum = sum + nums[i]
		} else {
			sum = nums[i]
		}
		ret = max(ret, sum)
	}
	return ret
}

// 解法三：分而治之，数组按中心点分成两个区间，最大子数组要么出现在数组左边，要么是右边，或者横跨这两个区间
// note：https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-cshi-xian-si-chong-jie-fa-bao-li-f/
func maxSubArray2(nums []int) int {
	return maxSubArrayDivide(nums, 0, len(nums)-1)
}
func maxSubArrayDivide(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}
	mid := (right-left)/2 + left
	leftMaxSum := maxSubArrayDivide(nums, left, mid)
	rightMaxSum := maxSubArrayDivide(nums, mid+1, right)

	leftCrossMaxSum := math.MinInt64
	leftCrossSum := 0
	for i := mid; i >= left; i-- {
		leftCrossSum += nums[i]
		leftCrossMaxSum = max(leftCrossMaxSum, leftCrossSum)
	}

	rightCrossMaxSum := math.MinInt64
	rightCrossSum := 0
	for i := mid + 1; i <= right; i++ {
		rightCrossSum += nums[i]
		rightCrossMaxSum = max(rightCrossMaxSum, rightCrossSum)
	}

	crossMaxSum := leftCrossMaxSum + rightCrossMaxSum

	return max(max(leftMaxSum, rightMaxSum), crossMaxSum)
}
