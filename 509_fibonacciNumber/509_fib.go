package main

// 递归算法的时间复杂度计算：子问题个数乘以解决一个子问题需要的时间

// 解法一：暴力递归（自顶向下的解法）。时间复杂度为O(2^n)。
// note：存在很多的重叠子问题，导致大量的重复计算。
func fib0(n int) int {
	if n <= 1 {
		return n
	}
	return fib0(n-1) + fib0(n-2)
}

// 解法二：带备忘录的递归（自顶向下的解法）。时间和空间复杂度都为O(n)。
// note：用一个备忘录来存放曾经计算过的子问题的结果，显然就不需要重复计算某些子问题了，所以需要计算的子问题个数为n（将递归树进行“剪枝“）
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	// 这里要初始化一个n+1长度的数组，因为这里要记录到n的值
	dp := make([]int, n+1)
	return fib1Core(dp, n)
}
func fib1Core(dp []int, n int) int {
	// 递归结束条件
	if n <= 1 {
		return n
	}
	// 等于0表示这个子问题还没有计算，则按照题目给出的公式进行递归计算，
	// 计算完的结果需要存起来，然后返回
	if dp[n] != 0 {
		return dp[n]
	}
	// 已经计算过就直接返回即可
	dp[n] = fib1Core(dp, n-1) + fib1Core(dp, n-2)
	return dp[n]
}

// 解法三：带备忘录的迭代（自底向上的解法）。时间复杂度为O(n)
func fib2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	// base case
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 解法四：不带备忘录的迭代（自底向上的解法）。时间复杂度为O(n)
// note：这里其实不需要dp table，因为我们可以看到n的结果只与前面两个的结果相关联
func fib3(n int) int {
	if n <= 1 {
		return n
	}
	// base case
	base0 := 0
	base1 := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = base0 + base1
		base0 = base1
		base1 = result
	}
	return result
}