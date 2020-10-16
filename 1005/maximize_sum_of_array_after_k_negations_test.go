package main

import (
	"fmt"
	"testing"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	type in struct {
		A []int
		K int
	}

	tests := []struct {
		in  in
		out int
	}{
		{
			in{[]int{4, 2, 3}, 1},
			5,
		},
		{
			in{[]int{3, -1, 0, 2}, 3},
			6,
		},
		{
			in{[]int{2, -3, -1, 5, -4}, 2},
			13,
		},
		{
			in{[]int{-8, 3, -5, -3, -5, -2}, 6},
			22,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%v, %d", tt.in.A, tt.in.K), func(t *testing.T) {
			t.Parallel()
			if got := largestSumAfterKNegations(tt.in.A, tt.in.K); got != tt.out {
				t.Errorf("got %d, want %d", got, tt.out)
			}
		})
	}
}
