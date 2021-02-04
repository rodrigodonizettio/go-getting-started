package main

import (
	"fmt"

	"rodrigodonizettio.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world from Go!")
	fmt.Println(quote.Glass())

	message := greetings.Hail("donizetti")
	fmt.Println(message)
}
