package hello

import (
	"fmt"
)

const (
	french        = "French"
	spanish       = "Spanish"
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) string {
	switch language {
	case french:
		return frenchPrefix

	case spanish:
		return spanishPrefix
	}

	return englishPrefix
}

func main() {
	fmt.Println(Hello("Marouane", "french"))
}
