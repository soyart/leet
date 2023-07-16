package main

// Naive solution to:
// https://leetcode.com/problems/median-of-two-sorted-arrays
func findMedianSortedArrays(arr1 []int, arr2 []int) float64 {
	return med(naiveMerge(arr1, arr2))
}

// Merge sorted arrays `arr1` and `arr2` up to the median position(s)
func naiveMerge(arr1 []int, arr2 []int) []int {
	l1 := len(arr1)
	l2 := len(arr2)

	mergedLen := l1 + l2
	medianPos := (mergedLen + 1) / 2

	if mergedLen == 0 {
		panic("both arrays are empty")
	}

	pos := 0
	pos1 := 0
	pos2 := 0

	finished := false
	merged := make([]int, mergedLen)

	for {
		if pos > medianPos || finished {
			break
		}

		var r1, r2 *int
		if pos1 < l1 {
			r1 = &arr1[pos1]
		}
		if pos2 < l2 {
			r2 = &arr2[pos2]
		}

		switch {
		case r1 == nil && r2 == nil:
			finished = true

		case r1 != nil && r2 != nil:
			n1, n2 := *r1, *r2

			switch {
			case n1 < n2:
				merged[pos] = n1
				pos1++

			default:
				merged[pos] = n2
				pos2++
			}

		case r1 != nil:
			merged[pos] = *r1
			pos1++

		case r2 != nil:
			merged[pos] = *r2
			pos2++
		}

		pos++
	}

	return merged
}

func med(arr []int) float64 {
	l := len(arr)
	if l == 0 {
		panic("empty array")
	}

	mid := l / 2

	if l%2 == 0 {
		return (float64(arr[mid]) + float64(arr[mid-1])) / 2
	}

	return float64(arr[mid])
}
