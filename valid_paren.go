package main

func isValid(s string) bool {
	var opens, closes []rune

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			opens = append(opens, char)

		case ')', '}', ']':
			if len(opens) < 1 {
				return false
			}

			closes = append(closes, char)
		}
	}

	if len(opens) != len(closes) {
		return false
	}

	if len(opens) == 0 {
		return true
	}

	// ()[]{}
	// opens:  ( [ {
	// closes: ) ] }

	// [{(foo)}]
	// opens: [, {, (
	// closes: ), }, ]

	for i, j := 0, len(opens)-1; i <= j; i, j = i+1, j-1 {
		open, close := opens[i], closes[j]

		switch open {
		case '(':
			if close != ')' {
				return false
			}
		case '{':
			if close != '}' {
				return false
			}
		case '[':
			if close != ']' {
				return false
			}
		}
	}

	return true
}
