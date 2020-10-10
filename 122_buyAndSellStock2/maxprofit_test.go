package main

import "testing"

func TestMaxProfit(t *testing.T) {
	var testSuite = []struct {
		prices         []int
		expectedProfit int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, expectedProfit: 7},
		{prices: []int{7, 6, 4, 3, 1}, expectedProfit: 0},
		{prices: []int{1, 2, 3, 4, 5}, expectedProfit: 4},
	}
	for i := 0; i < len(testSuite); i++ {
		got0 := maxProfit0(testSuite[i].prices)
		if got0 != testSuite[i].expectedProfit {
			t.Fatalf("maxProfit0: want: %d, got: %d\n", testSuite[i].expectedProfit, got0)
		}
		//got1 := maxProfit1(testSuite[i].prices)
		//if got1 != testSuite[i].expectedProfit {
		//	t.Fatalf("maxProfit1: want: %d, got: %d\n", testSuite[i].expectedProfit, got1)
		//}
		//got2 := maxProfit2(testSuite[i].prices)
		//if got2 != testSuite[i].expectedProfit {
		//	t.Fatalf("maxProfit2: want: %d, got: %d\n", testSuite[i].expectedProfit, got2)
		//}
	}
}
