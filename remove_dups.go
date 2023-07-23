package main

// https://leetcode.com/problems/remove-duplicates-from-sorted-array
// Note: nums is sorted
func removeDuplicates(nums []int) int {
	if l := len(nums); l <= 1 {
		return l
	}

	// k is the number of uniques, p is insertion pointer
	k, p := 1, 1
	prev := nums[0]

	for _, n := range nums[1:] {
		if n == prev {
			continue
		}

		// New unique values
		prev = n
		nums[p] = n

		k, p = k+1, p+1
	}

	return k
}
