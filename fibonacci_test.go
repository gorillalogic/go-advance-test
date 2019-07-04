// Copyright 2019 Gorilla Logic, Inc.
// All rights reserved.
package fib

import "testing"

func TestFibonacci(t *testing.T) {
	data := []struct {
		n    uint
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296}, {90, 2880067194370816120},
	}

	for _, d := range data {
		if got := Fibonacci(d.n); got != d.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}