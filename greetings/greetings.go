package greetings

import "fmt"

func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	message, error := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, error
}
