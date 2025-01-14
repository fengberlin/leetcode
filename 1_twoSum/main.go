package main

import "fmt"

// https://leetcode.cn/problems/two-sum/description/
func main() {
	nums := []int{2,7,11,15}
	target := 9
	i := twoSum(nums, target)
	fmt.Println(i)
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
    m := make(map[int]int)
	for i := range nums {
		if idx, ok := m[target-nums[i]]; ok {
			return []int{idx, i}
		}
		m[nums[i]] = i
	}
	return nil
}