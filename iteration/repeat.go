package iteration

func Repeat(character string, repeatCount int) string {
	if repeatCount < 0 {
		repeatCount = 0
	}

	var repeatedChar string

	for i := 0; i < repeatCount; i++ {
		repeatedChar += character
	}
	return repeatedChar
}
