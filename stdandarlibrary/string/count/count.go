package count

import "strings"

func CountMany(word, subWord string) int {
	return strings.Count(word, subWord)
}
