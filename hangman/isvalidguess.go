package hangman

func IsValidGuess(n interface{}) bool {
	var q bool = true
	switch n := n.(type) {
	case string:
		for _, i := range n {
			if i == '\n' {
				q = true
			} else if i != ' ' && i < 'A' || i > 'Z' && i < 'a' || i > 'z' {
				q = false
				return q
			}
		}
	case rune:
		if n == '\n' {
			q = true
		} else if n != ' ' && n < 'A' || n > 'Z' && n < 'a' || n > 'z' {
			q = false
			return q
		}
	}
	return q
}
