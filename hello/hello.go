package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"AAA", "BBB", "CCC"}

	//request a greeting message
	message, err := greetings.Hellos(names)

	//if an error was returned, print it to the console and exit program
	if err != nil {
		log.Fatal(err)
	}
	// if no error
	fmt.Println(message)
}
