package containsany

import "strings"

func WordContainsAny(word, subWord string) bool {
	return strings.ContainsAny(word, subWord)
}
