package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type Test struct {
		nums         []int
		expectedK    int
		expectedNums []int
	}

	tests := []Test{
		{
			nums:         []int{1, 1, 2},
			expectedK:    2,
			expectedNums: []int{1, 2},
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedK:    5,
			expectedNums: []int{0, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		k := removeDuplicates(test.nums)

		for i, n := range test.nums[:k] {
			if e := test.expectedNums[i]; e != n {
				t.Log(test.nums)
				t.Fatalf("unexpected value: expecting %d, got %d", e, n)
			}
		}
	}
}
