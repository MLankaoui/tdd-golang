package iteration

func Repeat(character string) string {
	var repeatedChar string

	for i := 0; i < 5; i++ {
		repeatedChar += character
	}
	return repeatedChar
}
