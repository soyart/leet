package main

import "testing"

func TestNaiveMergedMedian(t *testing.T) {
	tests := map[float64][2][]int{
		1: {
			{1},
			{},
		},
		2: {
			{1, 2},
			{3},
		},
		2.5: {
			{},
			{2, 3},
		},
		// 2.5: {
		// 	{1, 2},
		// 	{3, 4},
		// },
		5: {
			{1, 2, 3, 4},
			{5, 6, 7, 8, 9},
		},
		4: {
			{2, 2, 3, 4, 5, 6, 10, 12},
			{1, 2, 7},
		},
	}

	for expected, inputs := range tests {
		merged := naiveMerge(inputs[0], inputs[1])
		median := med(merged)

		if median != expected {
			t.Log("inputs", inputs)
			t.Log("merged", merged)

			t.Fatalf("unexpected median: expected %f, got %f", expected, median)
		}
	}
}
