package main

func isValid(s string) bool {
	var openStack []rune

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			openStack = append(openStack, char)

		case ')', '}', ']':
			l := len(openStack)
			if l < 1 {
				return false
			}

			lastOpen := openStack[l-1]

			switch char {
			case ')':
				if lastOpen != '(' {
					return false
				}
			case '}':
				if lastOpen != '{' {
					return false
				}
			case ']':
				if lastOpen != '[' {
					return false
				}
			}

			if l > 0 {
				openStack = openStack[:l-1]
			}
		}
	}

	if len(openStack) != 0 {
		return false
	}

	return true
}
