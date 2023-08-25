package main

import "fmt"

const englishHelloPrefix = "Hello, "

func SayHello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(SayHello("X"))
}
