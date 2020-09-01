package main

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	var testSuite = []struct {
		coins  []int
		amount int
		want   int
	}{
		{coins: []int{1, 2, 5}, amount: 11, want: 3},
		{coins: []int{2}, amount: 3, want: -1},
		{coins: []int{1, 2147483647}, amount: 2, want: 2},
	}

	for i := 0; i < len(testSuite); i++ {
		got0 := coinChange0(testSuite[i].coins, testSuite[i].amount)
		got1 := coinChange1(testSuite[i].coins, testSuite[i].amount)
		got2 := coinChange2(testSuite[i].coins, testSuite[i].amount)
		got3 := coinChange3(testSuite[i].coins, testSuite[i].amount)
		got4 := coinChange4(testSuite[i].coins, testSuite[i].amount)
		if got0 != testSuite[i].want {
			t.Fatalf("coinChange0: want: %d, got: %d\n", testSuite[i].want, got0)
		}
		if got1 != testSuite[i].want {
			t.Fatalf("coinChange1: want: %d, got: %d\n", testSuite[i].want, got1)
		}
		if got2 != testSuite[i].want {
			t.Fatalf("coinChange2: want: %d, got: %d\n", testSuite[i].want, got2)
		}
		if got3 != testSuite[i].want {
			t.Fatalf("coinChange3: want: %d, got: %d\n", testSuite[i].want, got3)
		}
		if got4 != testSuite[i].want {
			t.Fatalf("coinChange4: want: %d, got: %d\n", testSuite[i].want, got4)
		}
	}
}
