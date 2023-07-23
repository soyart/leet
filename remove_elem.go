package main

// https://leetcode.com/problems/remove-element/
// Caveat: Before returning k, the first k elements of nums must only contain
// values without val.
// The length and order of the fist 5 k elements are not important.
func removeElement(nums []int, val int) int {
	// k is the number of all elements that are not val
	// i is the insertion pointer for in-place removal
	var k, i int

	for _, n := range nums {
		if n != val {
			// Remove ith element in-place with n (which we know is not val)
			nums[i] = n

			// Increment both counters
			k, i = k+1, i+1
		}
	}

	return k
}
