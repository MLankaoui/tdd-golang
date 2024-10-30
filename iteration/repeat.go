package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeatedChar string

	for i := 0; i < repeatCount; i++ {
		repeatedChar += character
	}
	return repeatedChar
}
