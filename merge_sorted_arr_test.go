package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	type Test struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}

	tests := []Test{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums1:    []int{4, 0, 0, 0, 0, 0},
			m:        1,
			nums2:    []int{1, 2, 3, 5, 6},
			n:        5,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums1:    []int{1, 2, 4, 5, 6, 0},
			m:        5,
			nums2:    []int{3},
			n:        1,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums1:    []int{0, 0, 3, 0, 0, 0, 0, 0, 0},
			m:        3,
			nums2:    []int{-1, 1, 1, 1, 2, 3},
			n:        6,
			expected: []int{-1, 0, 0, 1, 1, 1, 2, 3, 3},
		},
	}

	for _, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)

		if reflect.DeepEqual(test.nums1, test.expected) {
			continue
		}

		t.Fatalf("unexpected value, expecting %v, got %v", test.expected, test.nums1)
	}
}
