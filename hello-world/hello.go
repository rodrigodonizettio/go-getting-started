package main

import (
	"fmt"
	"log"

	"rodrigodonizettio.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world from Go!")
	fmt.Println(quote.Glass())

	//Setting properties to the Logger. "0" disables the Timestamp and the Source File Name
	log.SetPrefix("Log from greetings.go: ")
	log.SetFlags(0)

	fmt.Println(message)
}
