package parenthesis

func isValid(s string) bool {
	pairs := make([]rune, 0)
	for _, c := range s {
		switch c {
		case ')':
			if len(pairs) == 0 || pairs[len(pairs)-1] != '(' {
				return false
			}
			pairs = pairs[:len(pairs)-1]
		case ']':
			if len(pairs) == 0 || pairs[len(pairs)-1] != '[' {
				return false
			}
			pairs = pairs[:len(pairs)-1]
		case '}':
			if len(pairs) == 0 || pairs[len(pairs)-1] != '{' {
				return false
			}
			pairs = pairs[:len(pairs)-1]
		default:
			pairs = append(pairs, c)
		}
	}
	return len(pairs) <= 0
}
