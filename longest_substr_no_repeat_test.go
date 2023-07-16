package main

import "testing"

func TestLengthOfLongesrSubstring(t *testing.T) {
	tests := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"au":       2,
	}

	for s, expected := range tests {
		result := lengthOfLongestSubstring(s)
		if result != expected {
			t.Logf("input: %s", s)
			t.Fatalf("unexpected value: expecting %d, got %d", expected, result)
		}
	}
}

func TestLongestNoRepeat(t *testing.T) {
	tests := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   2,
		"au":       2,
	}

	for s, expected := range tests {
		result := longestNoRepeat(s)
		if result != expected {
			t.Logf("input: %s", s)
			t.Fatalf("unexpected value: expecting %d, got %d", expected, result)
		}
	}
}
