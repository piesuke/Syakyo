package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "

const spanishHelloPrefix = "Hola, "

const frenchHelloPrefix = "Bonjour, "

const spanish = "Spanish"

const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// (prefix string) は named return value。関数にある変数でreturnされるもの
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("world", ""))
}
