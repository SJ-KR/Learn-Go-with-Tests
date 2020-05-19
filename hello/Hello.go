package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHello = "Hello"
const spanishHello = "Hola"
const frenchHello = "Bonjour"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case spanish:
		return fmt.Sprintf("%s, %s", spanishHello, name)
	case french:
		return fmt.Sprintf("%s, %s", frenchHello, name)
	}
	return fmt.Sprintf("%s, %s", englishHello, name)
}
func main() {
	fmt.Println(Hello("World", "Spanish"))
	fmt.Println(Hello("World", "French"))
	fmt.Println(Hello("World", ""))
}
