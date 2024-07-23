package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//request a greeting message
	message, err := greetings.Hello("")
	//if an error was returned, print it to the console and exit program
	if err != nil {
		log.Fatal(err)
	}
	// if no error
	fmt.Println(message)
}
