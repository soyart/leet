package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := map[*[]int]int{
		{7, 1, 5, 3, 6, 4}: 5,
		{10, 20, 11, 29}:   19,
	}

	for input, expected := range tests {
		prices := *input
		result := maxProfit(prices)

		if result != expected {
			t.Logf("prices: %v", prices)
			t.Fatalf("unexpected value: expecting %d, got %d", expected, result)
		}
	}
}
