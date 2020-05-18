package main

import "fmt"

const eng = "Hello"

func Hello(name string) string {
	return fmt.Sprintf("%s, %s", eng, name)
}
func main() {
	fmt.Println(Hello("world"))
}
