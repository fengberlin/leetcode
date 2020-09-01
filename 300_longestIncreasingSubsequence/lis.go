package main

import "encoding/binary"

// note：dp[i]定义为以nums[i]这个数结尾的最长递增子序列的长度。
// 时间复杂度：O(n^2)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	binary.Varint()
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		// base case，最少也要包含一个数
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// 递增序列
			// nums[i]必然要大于nums[j]，才能将nums[i]nums[j]后面以形成更长的上升子序列。
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ret := 0
	for i := 0; i < len(dp); i++ {
		ret = max(ret, dp[i])
	}
	//ret := 0
	//for i := 0; i < len(nums); i++ {
	//	// base case，最少也要包含一个数
	//	dp[i] = 1
	//	for j := 0; j < i; j++ {
	//		// 递增序列
	//		// nums[i]必然要大于nums[j]，才能将nums[i]nums[j]后面以形成更长的上升子序列。
	//		if nums[i] > nums[j] {
	//			dp[i] = max(dp[i], dp[j]+1)
	//		}
	//	}
	//	ret = ret = max(ret, dp[i])
	//}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}