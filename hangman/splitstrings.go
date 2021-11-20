package hangman

func SplitStrings(str string) []string {
	k := 0
	no_white_spaces := false
	for index := range str {
		if no_white_spaces && index != 0 && (str[index-1] == '\n' || str[index-1] == '\t') && str[index] != '\n' && str[index] != '\t' {
			k++
		}
		if str[index] != '\n' && str[index] != '\t' {
			no_white_spaces = true
		}
	}
	k++

	x := 0
	ans := make([]string, k)
	ok := true
	for _, c := range str {
		if c == '\n' || c == '\t' {
			if !ok {
				x++
			}
			ok = true
			continue
		}
		ans[x] = ans[x] + string(c)
		ok = false
	}
	return ans
}
