package main

import "fmt"

// m = len(nums1), n = len(nums2)
// Store the result in num1
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Forward counters for nums1, nums2
	mergedLen := m + n

	// c is the insert pointer
	// c1 points to current element of nums1
	// c2 points to current element of nums2
	// temp points to values moved from nums1 to make way
	c, c1, c2, temp := 0, 0, 0, m

	for {
		if c == mergedLen {
			break
		}

		fmt.Println("nums1", nums1, "c", c, "c1", c1, "c2", c2)

		var p1, p2 *int

		if c1 < m {
			p1 = &nums1[c1]
		}

		if c2 < n {
			p2 = &nums2[c2]
		}

		switch {
		case p1 != nil && p2 != nil:
			n1, n2 := *p1, *p2

			if n1 <= n2 {
				c1++
			} else { // n2 < n1
				nums1[temp], nums1[c] = n1, n2
				c1, c2, temp = c1+1, c2+1, temp+1
			}

		// Only nums2 is left
		case p2 != nil:
			n2 := *p2
			occupant := nums1[c]

			if n2 < occupant || c >= temp {
				nums1[temp], nums1[c] = occupant, n2

				c2, temp = c2+1, temp+1
			}

			fmt.Println("eiei", nums1, "temp", temp, "c", c, "c2", c2, "n2", n2, "occupant", occupant)

		default:
			shifted := nums1[temp-1]
			current := nums1[c]

			if shifted < current {
				nums1[c] = shifted
				nums1[temp-1] = current
			}
		}

		c++
	}
}
