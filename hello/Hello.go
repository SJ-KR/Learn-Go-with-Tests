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

	return fmt.Sprintf("%s, %s", prefix(language), name)

}
func prefix(language string) (hello string) {
	switch language {
	case spanish:
		hello = spanishHello
	case french:
		hello = frenchHello
	default:
		hello = englishHello
	}
	return
}
func main() {
	fmt.Println(Hello("World", "Spanish"))
	fmt.Println(Hello("World", "French"))
	fmt.Println(Hello("World", ""))
}
