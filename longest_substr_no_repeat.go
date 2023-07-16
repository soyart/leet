package main

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	longest := 0

	for i := range s {
		length := 0
		seen := make(map[byte]struct{})

		for j := i; j < l; j++ {
			char := s[j]
			if _, found := seen[char]; found {
				break
			}

			length++
			seen[char] = struct{}{}
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}

// Similar to above, but from start, not for substrings
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
