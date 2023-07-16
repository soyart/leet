package main

func lengthOfLongestSubstring(s string) int {
	l := len(s)

	switch l {
	case 0, 1:
		return l
	}

	longest := 1
	for i := range s {
		for j := i; j <= l; j++ {
			subStr := s[i:j]
			if len(subStr) <= 1 {
				continue
			}

			count := longestNoRepeat(subStr)
			if count > longest {
				longest = count
			}
		}
	}

	return longest
}

func longestNoRepeat(s string) int {
	l := len(s)

	switch l {
	case 0, 1:
		return l
	}

	seen := make(map[rune]struct{})
	for i, char := range s {
		if _, ok := seen[char]; ok {
			return i
		}

		seen[char] = struct{}{}
	}

	return l
}
