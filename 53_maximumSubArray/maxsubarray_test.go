package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	var testSuite = []struct {
		nums           []int
		expectedMaxSum int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expectedMaxSum: 6},
		{nums: []int{-12, 11, -3, 4, -11, 2, 15, -5, 4}, expectedMaxSum: 18},
		{nums: []int{1, 13, -31, 12, -64, 28, 15, 5, -2}, expectedMaxSum: 48},
	}
	for i := 0; i < len(testSuite); i++ {
		got0 := maxSubArray0(testSuite[i].nums)
		if got0 != testSuite[i].expectedMaxSum {
			t.Fatalf("maxSubArray0: want: %d, got: %d\n", testSuite[i].expectedMaxSum, got0)
		}
		got1 := maxSubArray1(testSuite[i].nums)
		if got1 != testSuite[i].expectedMaxSum {
			t.Fatalf("maxSubArray1: want: %d, got: %d\n", testSuite[i].expectedMaxSum, got1)
		}
		got2 := maxSubArray2(testSuite[i].nums)
		if got2 != testSuite[i].expectedMaxSum {
			t.Fatalf("maxSubArray2: want: %d, got: %d\n", testSuite[i].expectedMaxSum, got2)
		}
	}
}
