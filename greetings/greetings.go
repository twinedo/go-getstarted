package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// if no name was given, return error message
	if name == "" {
		return "", errors.New("Empty name")
	}

	//if name was received, return value that eembeds the name in greeting message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
