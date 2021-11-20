package hangman

func Tried(s string) string {
	for _, i := range s {
		if i == '\n' {
			s = s[:len(s)-1] + ", "
		}
	}
	return s
}
