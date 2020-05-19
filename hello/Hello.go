package main

import "fmt"

const spanish = "Spanish"
const englishhello = "Hello"
const spanishhello = "Hola"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return fmt.Sprintf("%s, %s", spanishhello, name)
	}

	return fmt.Sprintf("%s, %s", englishhello, name)
}
func main() {
	fmt.Println(Hello("World", "Hello"))
}
