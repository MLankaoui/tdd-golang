package containsfunc

import "strings"

func ContainsFuncRune(word string) bool {
	return strings.ContainsFunc(word, F)
}

func F(letter rune) bool {
	return letter == 'a' || letter == 'h'
}
