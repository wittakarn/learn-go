package helloworld

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func SayHello(name string, language string) string {
	prefix := englishHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}

	if name == "" {
		return prefix + "World"
	}
	return prefix + name
}

func main() {
	fmt.Println(SayHello("Mr.X", "English"))
}
