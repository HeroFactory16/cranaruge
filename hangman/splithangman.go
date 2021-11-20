package hangman

import (
	"strings"
)

func SplitHangman(str string) []string {
	figure := strings.Split(str, "\n")
	hung := []string{}
	var tmp string
	for c, i := range figure {
		tmp += i + "\n"
		if c%8 == 7 && c != 0 {
			hung = append(hung, tmp)
			tmp = ""
		}
	}
	return hung
}
