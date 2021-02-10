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

	message, err := greetings.Hail("donizetti")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)


	//The variables "message" and "err" were already assigned in above statement
	message, err = greetings.Hail("")
	if err != nil {
		//IMPORTANT: If an error occurs the code execution is stopped, so the next lines won't be executed
		log.Fatal(err)
	}
	fmt.Println(message) //This line won't be executed if an error occurs (and it will)
}
