package main

import "testing"

func TestFib(t *testing.T) {
	var want = []struct {
		k int
		v int
	}{
		{k: 0, v: 0},
		{k: 1, v: 1},
		{k: 2, v: 1},
		{k: 3, v: 2},
		{k: 4, v: 3},
		{k: 5, v: 5},
		{k: 6, v: 8},
		{k: 7, v: 13},
		{k: 8, v: 21},
		{k: 9, v: 34},
	}

	for i := 0; i < len(want); i++ {
		got0 := fib0(want[i].k)
		got1 := fib1(want[i].k)
		got2 := fib2(want[i].k)
		got3 := fib3(want[i].k)
		if got0 != want[i].v {
			t.Fatalf("fib0: want: %d, got: %d\n", want[i].v, got0)
		}
		if got1 != want[i].v {
			t.Fatalf("fib1: want: %d, got: %d\n", want[i].v, got1)
		}
		if got2 != want[i].v {
			t.Fatalf("fib2: want: %d, got: %d\n", want[i].v, got2)
		}
		if got3 != want[i].v {
			t.Fatalf("fib3: want: %d, got: %d\n", want[i].v, got3)
		}
	}
}
