package main

import "testing"

func TestLengthOfLIS(t *testing.T) {
	var testSuite = []struct {
		nums        []int
		expectedLen int
	}{
		{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, expectedLen: 4},
		{nums: []int{5, 3, 108, -24, 33, -56, -1, 36}, expectedLen: 3},
	}
	for i := 0; i < len(testSuite); i++ {
		got := lengthOfLIS(testSuite[i].nums)
		if got != testSuite[i].expectedLen {
			t.Fatalf("lengthOfLIS: want: %d, got: %d\n", testSuite[i].expectedLen, got)
		}
	}
}
