package main

import (
	"fmt"
	"io"
)

func Countdown(w io.Writer) {
	fmt.Fprint(w, "3")
}

func main() {
	//Countdown()
}
