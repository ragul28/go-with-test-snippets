package main

import (
	"fmt"
)

const ENHelloPrefix = "Hello, "
const SPHelloPrefix = "Hola, "
const FNHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = FNHelloPrefix
	case "Spanish":
		prefix = SPHelloPrefix
	default:
		prefix = ENHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
