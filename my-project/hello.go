package main

import "fmt"

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("Tin", ""))
}
